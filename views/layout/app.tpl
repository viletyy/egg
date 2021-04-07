<!--
 * @Date: 2021-03-11 11:30:16
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:04:49
 * @FilePath: /egg/views/layout/app.tpl
-->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Egg</title>
  <link rel="stylesheet" href="/static/app.css">
</head>
<body>
<nav class="navbar has-shadow">
  <div class="container">
    <div class="navbar-brand">
        <a class="navbar-item" href="../"><img src="http://bulma.io/images/bulma-logo.png" alt="Bulma: a modern CSS framework based on Flexbox"/></a>
        <div class="navbar-burger burger" data-target="navMenu">
        </div>
    </div>
    <div class="navbar-menu" id="navMenu">
      <div class="navbar-end">
        <div class="navbar-item has-dropdown is-hoverable">
          {{if .IsLogin}}
          <a class="navbar-link">{{.User.Username}}</a>
          {{else}}
          <a class="navbar-link">未登陆</a>
          {{end}}

          <div class="navbar-dropdown">
            {{if .IsLogin}}
              <a class="navbar-item">profile</a>
              <hr class="navbar-divider" />
              <div class="navbar-item">
                <a href="{{urlfor "UserController.Loginout"}}">登出</a>
              </div>
            {{else}}
              <a class="navbar-item" href="{{urlfor "UserController.Register"}}">注册</a>
              <a class="navbar-item" href="{{urlfor "UserController.Login"}}">登陆</a>
            {{end}}
            {{if .IsAdmin}}
              <a class="navbar-item">后台管理</a>
            {{end}}
          </div>
        </div>
      </div>
    </div>
  </div>
</nav>
<section class="">
  {{if .flash.notice }}
  <div class="notification is-primary">
    <button class="delete"></button>
    <p>{{.flash.notice}}</p>
  </div>
  {{end}}
  {{if .flash.error }}
  <div class="notification is-danger">
    <button class="delete"></button>
    <p>{{.flash.error}}</p>
  </div>
  {{end}}
  {{if .flash.warning }}
  <div class="notification is-warning">
    <button class="delete"></button>
    <p>{{.flash.warning}}</p>
  </div>
  {{end}}
</section>
<section class="hero is-info">
<div class="hero-body">
    <div class="container">
        <div class="card">
            <div class="card-content">
                <div class="content">
                    <div class="control has-icons-left has-icons-right">
                        <input class="input is-large" type="search"/>
                        <span class="icon is-medium is-left">
                            <i class="fa fa-search"></i>
                        </span>
                        <span class="icon is-medium is-right">
                            <i class="fa fa-empire"></i>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</section>
<div class="box cta">
  <div class="columns is-mobile is-centered">
      <div class="field is-grouped is-grouped-multiline">
          <div class="control"><span class="tag is-link is-large">Link</span></div>
          <div class="control"><span class="tag is-success is-large">Success</span></div>
          <div class="control"><span class="tag is-black is-large">Black</span></div>
          <div class="control"><span class="tag is-warning is-large">Warning</span></div>
          <div class="control"><span class="tag is-danger is-large">Danger</span></div>
          <div class="control"><span class="tag is-info is-large">Info</span></div>
      </div>
  </div>
</div>
  <section class="container">
    <div class="level-item">
      <div class="is-multiline is-centered cards-container" id="sectioncontainer">
        {{.LayoutContent}}
      </div>
    </div>
  </section>
  <div class="columns is-mobile is-centered">
    <div class="column is-half is-narrow">
    </div>
  </div>
  <footer>
  <div class="box cta">
      <div class="columns is-mobile is-centered">
          <div class="field is-grouped is-grouped-multiline">
              <div class="control">
                  <div class="tags has-addons"><a class="tag is-link" href="https://bulmatemplates.github.io/bulma-templates/templates/kanban.html">Bulma Templates</a><span class="tag is-light">模版参考</span></div>
              </div>
              <div class="control">
                  <div class="tags has-addons"><a class="tag is-link" href="https://github.com/GolangMore/hello/">源代码</a><span class="tag is-light">hello  <i class="fa fa-github"></i></span></div>
              </div>
          </div>
      </div>
  </div>
  </footer>

  <script src="/static/app.js"></script>
</body>
</body>
</html>