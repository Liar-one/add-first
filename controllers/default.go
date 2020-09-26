 package controllers

import (
	"HelloBeego0406/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//name1:=c.GetString("name")
	//age1,err:=c.GetInt("age")


  name:=c.Ctx.Input.Query("name")
	age:=c.Ctx.Input.Query("age")
	sex:=c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//
	if name!="admit"||age!="18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	if name=="admit"&& age=="18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
		return
	}
	fmt.Println(name)
	fmt.Println(age)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"


}
 //Post请求
//func (c*MainController) Post() {
//	fmt.Println("post的请求")
//	user := c.Ctx.Request.FormValue("user")
//	fmt.Println("用户名是：", user)
//	psd := c.Ctx.Request.FormValue("psd")
//	fmt.Println("密码：", psd)
//
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//
//	fmt.Println("发送请求")
//
//	//与固定值进行比较
//	if user != "admin" || psd != "123456" {
//		//失败页面
//		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
//		return
//	}
//
//	//成功页面
//	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确")) * /
//}
  func(c *MainController)Post(){
  	dataByes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
  	if err !=nil{
  		c.Ctx.WriteString("数据失败")
  		return
	}
	//json包
	var person models.Person
  	err =json.Unmarshal(dataByes,&person)
  	if err !=nil{
  		c.Ctx.WriteString("数据解析失败")
  		return
	}
	fmt.Println("用户名",person.Name,"年龄",person.Age)
  	c.Ctx.WriteString("用户名："+person.Name)
	}

