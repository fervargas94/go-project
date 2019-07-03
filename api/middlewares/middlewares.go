package middleware

import (
	"bufio"
	"fmt"
	"github.com/fervargas94/proxy-app/api/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris"
)

// Queue is the struct to store information
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que stack declaration
var Que []string
var lowQueue, medQueue, highQueue []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

// Read Repository interface implementation
func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middlewares/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var final []*Queue
	tmp := &Queue{}
	count := 0
	for scanner.Scan() {
		count++
		if scanner.Text() == "" {
			count = 0
			continue
		}
		switch count {
		case 1:
			tmp.Domain = scanner.Text()
		case 2:
			r := strings.Split(scanner.Text(), ":")[1]
			res, _ := strconv.Atoi(r)
			tmp.Weight = res
		case 3:
			r := strings.Split(scanner.Text(), ":")[1]
			res, _ := strconv.Atoi(r)
			tmp.Priority = res
			// persist tmp struct
			final = append(final, tmp)
			// clean tmp struct
			tmp = &Queue{}
		}
	}
	return final
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}
	fmt.Println("FROM HEADER", domain)
	var priorityDefault int
	for _, row := range repo.Read() {
		if(row.Domain == domain){
			priorityDefault = utils.CalculatePriority(row.Priority, row.Weight)
		}
	}

	fmt.Println("PRIORITY DEFAULT", priorityDefault)
	fmt.Println("ANETES", lowQueue, medQueue, highQueue)

	if priorityDefault == 1 {
		lowQueue = append(lowQueue, domain)
	} else if priorityDefault == 2 {
		medQueue = append(medQueue, domain)
	} else if priorityDefault == 3 {
		highQueue = append(highQueue, domain)
	}

	fmt.Println("DESPUES", lowQueue, medQueue, highQueue)

	Que = nil

	Que = append(Que, lowQueue...)
	Que = append(Que, medQueue...)
	Que = append(Que, highQueue...)

	c.Next()
}