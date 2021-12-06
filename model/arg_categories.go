package model

import (
	"fmt"
	"strconv"
	"strings"
)

type CategoriesFlag []int

func (c *CategoriesFlag) String() string {
	return fmt.Sprint(*c)
}

func (c *CategoriesFlag) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		if v == "" {
			continue
		}
		if i, err := strconv.Atoi(v); err == nil {
			*c = append(*c, i)
		} else {
			return err
		}
	}

	return nil
}

func (c *CategoriesFlag) Contains(x int) bool {
	for _, v := range *c {
		if v == x {
			return true
		}
	}
	return false
}
