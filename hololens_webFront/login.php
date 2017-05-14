<?php

require_once './lib/header.php';

//Redirect to index if already logged in
if(is_logged_in()) {
    redirect('index.php');
}

if(is_post_request()) {
    $username = $_POST['username'];
    $password = $_POST['password'];
    $info = verify_user($username, $password);
    if(is_null($info) || ($info->Password !== $password)) {
        redirect('login.php?error=mismatch');
    } else {
        log_in($info->Username, $info->Email);
        redirect('index.php');
    }
}

?>
<!DOCTYPE html>
<html >
<head>
  <meta charset="UTF-8">
  <title>Sign-Up/Login Form</title>
  <link href='https://fonts.googleapis.com/css?family=Titillium+Web:400,300,600' rel='stylesheet' type='text/css'>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/5.0.0/normalize.min.css">

  
      <link rel="stylesheet" href="./css/login.css">

  
</head>

<body>
  <div class="form">
      
      <ul class="tab-group">
        <li class="tab"><a href="/signup.php">Sign Up</a></li>
        <li class="tab active"><a href="/login.php">Log In</a></li>
      </ul>

      <div id="login">
          <h1>Welcome Back!</h1>

          <form action="" method="post">

              <div class="field-wrap">
                  <label>
                      Username<span class="req">*</span>
                  </label>
                  <input type="text" name="username" required autocomplete="off"/>
              </div>

              <div class="field-wrap">
                  <label>
                      Password<span class="req">*</span>
                  </label>
                  <input type="password" name="password" required autocomplete="off"/>
              </div>

              <?php if(array_key_exists('error', $_GET)): ?>
                  <div class="alert alert-danger">
                      <?php
                      switch ($_GET['error']) {
                          case 'mismatch': echo 'Incorrect username/password combination'; break;
                          default: echo 'Unknown error occurred';
                      }
                      ?>
                  </div>
              <?php endif; ?>

              <!-- <p class="forgot"><a href="#">Forgot Password?</a></p> -->

              <button class="button button-block"/>Log In</button>

          </form>

      </div>
      
</div> <!-- /form -->
  <script src='https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js'></script>
  <script src="./js/login.js?v=2"></script>
</body>
</html>
