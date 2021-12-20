package main

// 一个部门下面可以存在子部门和员工, 员工下面不能再包含其他节点.
// 需要统计所有员工人数

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

// 实现 Iorganization 接口的 Count()
func (Employee) Count() int {
	return 1
}

type Department struct {
	Name string

	SubOrganizations []IOrganization
}

// 实现 Iorganization 接口的 Count()
func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

// 由于是接口 所以实现了接口的 Employee 和 Department 都可以进行append
func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	root.AddSub(&Employee{})
	root.AddSub(&Employee{})
	for i := 0; i < 10; i++ {
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}
