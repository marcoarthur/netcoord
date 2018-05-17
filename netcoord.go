package netcoord

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/serf/coordinate"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

/* Loggers */
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

/* Initialize Loggers */
func InitLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func GenerateRandomCoordinate() *coordinate.Coordinate {
	config := coordinate.DefaultConfig()
	coord := coordinate.NewCoordinate(config)
	for i := range coord.Vec {
		coord.Vec[i] = rand.NormFloat64()
	}
	coord.Error = rand.NormFloat64()
	coord.Adjustment = rand.NormFloat64()
	return coord
}

/* Questions:

1) How to set two nodes ? i  j
2) What coordinate does a node get by default?
3) How to simulate/aquire Rtt between i, j ?
4) How calculate update coordinate after Rtt (check on both theoritic and pratical) ?
5)

6) Is possible to run a centralizade algorithm ?
   "The centralized algorithm described in Section 2.3 computes co-
   ordinates for all nodes given all RTTs."
6)

*/

func CreateNode() *coordinate.Client {

	/* Create a New client usind default coordinates */
	cli, err := coordinate.NewClient(coordinate.DefaultConfig())

	if err != nil {
		Error.Printf("Error creating a new client: %s", err.Error())
	}

	return cli
}

func Walk() {
	InitLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("Understanding a Network Coordination")

	/* Create a New client usind default coordinates */
	n1 := CreateNode()

	Info.Println("Coordinates of n1")
	spew.Dump(n1.GetCoordinate())

	/* Showing coordinate vector */
	Info.Printf("Origin: %v", n1.GetCoordinate().Vec)

	/* Create another Client, with other coord */

	n2 := CreateNode()
	c1 := coordinate.NewCoordinate(coordinate.DefaultConfig())
	c1.Vec[0] = 1
	c1.Height = 0
	n2.SetCoordinate(c1)

	Info.Println("Coordinates of n2")
	spew.Dump(n2.GetCoordinate())

	/* Distances ?*/
}
