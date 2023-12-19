package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ToProcess struct {
	record   map[string]int
	workflow string
}

type ToProcessRange struct {
	recordRange map[string]Range
	workflow    string
}

type Rule struct {
	variable          string
	lessThan          bool
	compared          int
	resultantWorkflow string
	noCondition       bool
}

type Range struct {
	start int
	end   int
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n\n")
	// part 1
	workflows := map[string][]Rule{}
	for _, line := range strings.Split(lines[0], "\n") {
		parsedOutName := strings.Split(line, "{")
		name := parsedOutName[0]
		parsedOutName = strings.Split(parsedOutName[1], "}")
		ruleString := parsedOutName[0]
		rules := []Rule{}
		for _, rule := range strings.Split(ruleString, ",") {
			if strings.Contains(rule, ":") {
				splitRule := strings.Split(rule, ":")
				resultantWorkflow := splitRule[1]
				lessThan := true
				compared := 0
				variable := ""
				if strings.Contains(splitRule[0], "<") {
					split := strings.Split(splitRule[0], "<")
					variable = split[0]
					compared, _ = strconv.Atoi(split[1])
				} else {
					split := strings.Split(splitRule[0], ">")
					variable = split[0]
					compared, _ = strconv.Atoi(split[1])
					lessThan = false
				}
				rules = append(rules, Rule{variable: variable, lessThan: lessThan, compared: compared, resultantWorkflow: resultantWorkflow, noCondition: false})
			} else {
				rules = append(rules, Rule{resultantWorkflow: rule, noCondition: true})
			}
		}
		workflows[name] = rules
	}

	// process records
	processQueue := []ToProcess{}
	for _, line := range strings.Split(lines[1], "\n") {
		pattern := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
		record := map[string]int{}
		matches := pattern.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		m, _ := strconv.Atoi(matches[2])
		a, _ := strconv.Atoi(matches[3])
		s, _ := strconv.Atoi(matches[4])
		record["x"] = x
		record["m"] = m
		record["a"] = a
		record["s"] = s
		record["sum"] = x + m + a + s
		processQueue = append(processQueue, ToProcess{record: record, workflow: "in"})
	}

	accepted := []map[string]int{}
	for i := 0; i < len(processQueue); i++ {
		if processQueue[i].workflow == "A" {
			accepted = append(accepted, processQueue[i].record)
			continue
		}
		if processQueue[i].workflow == "R" {
			continue
		}
		workflow := workflows[processQueue[i].workflow]
		record := processQueue[i].record
		for _, rule := range workflow {
			if rule.noCondition == true {
				processQueue = append(processQueue, ToProcess{record: record, workflow: rule.resultantWorkflow})
			} else {
				lessThan := rule.lessThan
				if lessThan {
					variable := rule.variable
					compared := rule.compared
					if processQueue[i].record[variable] < compared {
						processQueue = append(processQueue, ToProcess{record: record, workflow: rule.resultantWorkflow})
						break
					}
				} else {
					variable := rule.variable
					compared := rule.compared
					if processQueue[i].record[variable] > compared {
						processQueue = append(processQueue, ToProcess{record: record, workflow: rule.resultantWorkflow})
						break
					}
				}

			}
		}
	}
	sum := 0
	for _, accept := range accepted {
		sum += accept["sum"]
	}
	fmt.Println(sum)

	// part 2
	RecordRange := map[string]Range{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}
	recordRangeQueue := []ToProcessRange{{recordRange: RecordRange, workflow: "in"}}
	acceptedRanges := []map[string]Range{}

	for i := 0; i < len(recordRangeQueue); i++ {
		if recordRangeQueue[i].workflow == "A" {
			acceptedRanges = append(acceptedRanges, recordRangeQueue[i].recordRange)
			continue
		}
		if recordRangeQueue[i].workflow == "R" {
			continue
		}
		workflow := workflows[recordRangeQueue[i].workflow]
		recordRange := recordRangeQueue[i].recordRange
		for _, rule := range workflow {
			if rule.noCondition == true {
				recordRangeQueue = append(recordRangeQueue, ToProcessRange{recordRange: recordRange, workflow: rule.resultantWorkflow})
			} else {
				lessThan := rule.lessThan
				if lessThan {
					variable := rule.variable
					compared := rule.compared
					if recordRange[variable].end < compared {
						recordRangeQueue = append(recordRangeQueue, ToProcessRange{recordRange: recordRange, workflow: rule.resultantWorkflow})
						break
					}
					if compared > recordRange[variable].start && compared < recordRange[variable].end {
						lessThanRange := map[string]Range{"x": recordRange["x"], "m": recordRange["m"], "a": recordRange["a"], "s": recordRange["s"]}
						lessThanRange[variable] = Range{recordRange[variable].start, compared - 1}
						recordRange[variable] = Range{compared, recordRange[variable].end}
						recordRangeQueue = append(recordRangeQueue, ToProcessRange{recordRange: lessThanRange, workflow: rule.resultantWorkflow})
					}
				} else {
					variable := rule.variable
					compared := rule.compared
					if recordRange[variable].start > compared {
						recordRangeQueue = append(recordRangeQueue, ToProcessRange{recordRange: recordRange, workflow: rule.resultantWorkflow})
						break
					}
					if compared > recordRange[variable].start && compared < recordRange[variable].end {
						greaterThanRange := map[string]Range{"x": recordRange["x"], "m": recordRange["m"], "a": recordRange["a"], "s": recordRange["s"]}
						greaterThanRange[variable] = Range{compared + 1, recordRange[variable].end}
						recordRange[variable] = Range{recordRange[variable].start, compared}
						recordRangeQueue = append(recordRangeQueue, ToProcessRange{recordRange: greaterThanRange, workflow: rule.resultantWorkflow})
					}
				}
			}
		}
	}
	sum = 0
	for _, recordRange := range acceptedRanges {
		sum += (recordRange["x"].end - recordRange["x"].start + 1) * (recordRange["m"].end - recordRange["m"].start + 1) * (recordRange["a"].end - recordRange["a"].start + 1) * (recordRange["s"].end - recordRange["s"].start + 1)
	}
	fmt.Println(sum)
}
