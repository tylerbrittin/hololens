<?php



require_once __DIR__.'/lib/session.php';
require_once __DIR__.'/lib/util.php';
require_once __DIR__.'/lib/api.php';


//Redirect to index if already logged in
if(is_logged_in()) {
    redirect('/index.php');
}

if(is_post_request()) {

    $first_name = $_POST['first_name'];
    $last_name = $_POST['last_name'];
    $username = $_POST['username'];
    $email = $_POST['email'];
    $password = $_POST['password'];

    if(!check_username_availability($username)) {
        redirect('/signup.php?error=unavailable');
    }

    $success = create_user([
        'first_name' => $first_name,
        'last_name' => $last_name,
        'username' => $username,
        'email' => $email,
        'password' => $password
    ]);

    if(!$success) {
        redirect('/signup.php?error=unknown');
    }

    log_in($username, $email);
    redirect('/index.php');

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
        <li class="tab active"><a href="/signup.php">Sign Up</a></li>
        <li class="tab"><a href="/login.php">Log In</a></li>
      </ul>

      <div id="signup">
          <h1>Sign Up!</h1>

          <form action="" method="post">

              <div class="top-row">
                  <div class="field-wrap">
                      <label>
                          First Name<span class="req">*</span>
                      </label>
                      <input type="text" name="first_name" required />
                  </div>

                  <div class="field-wrap">
                      <label>
                          Last Name<span class="req">*</span>
                      </label>
                      <input type="text" name="last_name" required />
                  </div>
              </div>

              <div class="field-wrap">
                  <label>
                      Username<span class="req">*</span>
                  </label>
                  <input type="text" name="username" required autocomplete="off"/>
              </div>

              <div class="field-wrap">
                  <label>
                      Email<span class="req">*</span>
                  </label>
                  <input type="email" name="email" required/>
              </div>

              <div class="field-wrap">
                  <label>
                      Create Password<span class="req">*</span>
                  </label>
                  <input type="password" name="password" required autocomplete="off"/>
              </div>

              <?php if(array_key_exists('error', $_GET)): ?>
                  <div class="alert alert-danger">
                      <?php
                      switch ($_GET['error']) {
                          case 'unavailable': echo 'The username you chose has already been taken, please choose another'; break;
                          default: echo 'Unknown error occurred';
                      }
                      ?>
                  </div>
              <?php endif; ?>

              <button type="submit" class="button button-block"/>Get Started</button>

          </form>

      </div>
      
</div> <!-- /form -->
  <script src='https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js'></script>
    <script src="./js/login.js?v=8"></script>
    <script>
      $(document).ready(function() {
        $('.form').on('submit', function(e) {

          var has_errors = false;

          if (!$('input[name="first_name"]').val()) {
            alert('First name is required');
            has_errors = true;
          }

          if (!$('input[name="last_name"]').val() && !has_errors) {
            alert('Last name is required');
            has_errors = true;
          }

          if (!$('input[name="username"]').val() && !has_errors) {
            alert('Username is required');
            has_errors = true;
          }

          if (!$('input[name="email"]').val() && !has_errors) {
            alert('Email is required');
            has_errors = true;
          } else {
            if (!validateEmail($('input[name="email"]').val()) && !has_errors) {
              alert('Please enter a valid email address');
              has_errors = true;
            }
          }

          if (!$('input[name="password"]').val() && !has_errors) {
            alert('Password is required');
            has_errors = true;
          }

          var password = $('input[name="password"]').val();

          if (password.length < 6 && !has_errors) {
            alert('Password must be 6 characters or long');
            has_errors = true;
          }

          var regex = /^(?=.*[A-Z]).+$/;

          if(!regex.test(password) && !has_errors) {
            alert('Password must contain at least 1 uppercase character');
            has_errors = true;
          }

          var regex = /^(?=.*[0-9]).+$/;

          if(!regex.test(password) && !has_errors) {
            alert('Password must contain at least 1 number');
            has_errors = true;
          } 

          if (has_errors) {
            e.preventDefault();
          }
        });

      })
      
function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
}
    </script>
</body>
</html>
