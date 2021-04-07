<!--
 * @Date: 2021-03-12 14:48:43
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-13 19:03:24
 * @FilePath: /hello/views/usercontroller/login.tpl
-->

<!--面包屑-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
    <li class="is-active"><a href="#" aria-current="page">登陆</a></li>
  </ul>
</nav>

<form action="{{urlfor "UserController.LoginHandle"}}" class="box" method="post">
  {{ .xsrfdata }}
  <div class="field">
    <label class="label" for="username">密码：</label>
    <input class="input" type="text" name="username">
  </div>
  <div class="field">
    <label class="label" for="password">密码：</label>
    <input class="input" type="password" name="password">
  </div>
  <button class="button is-primary" type="submit">登陆</button>
  <button class="button is-default" type="reset">重置</button>
</form>