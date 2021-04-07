<!--
 * @Date: 2021-03-12 16:03:03
 * @LastEditors: viletyy
 * @LastEditTime: 2021-03-12 18:51:58
 * @FilePath: /hello/views/usercontroller/register.tpl
-->

<!--面包屑-->
<nav class="breadcrumb" aria-label="breadcrumbs">
  <ul>
    <li><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
    <li class="is-active"><a href="#" aria-current="page">注册</a></li>
  </ul>
</nav>

<form action="{{urlfor "UserController.RegisterHandle"}}" method="post">
  {{ .xsrfdata }}
  <div class="field">
    <label class="label" for="username">用户名：</label>
    <input class="input" type="text" name="username">
  </div>
  <div class="field">
    <label class="label" for="password">密码：</label>
    <input class="input" type="password" name="password">
  </div>
  <div class="field">
    <label class="label" for="password_confirmation">密码确认：</label>
    <input class="input" type="password" name="password_confirmation">
  </div>
  <button class="button is-primary" type="submit">注册</button>
  <button class="button is-default" type="reset">重置</button>
</form>