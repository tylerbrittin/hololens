<?php

require_once './lib/header.php';

//Redirect to login if not logged in (login is required)
if(!is_logged_in()) {
    redirect('/login.php');
}

$user = get_logged_in_user();

?>
<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Nebula Shop</title>

    <!-- Bootstrap Core CSS -->
    <link href="vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">

    <!-- Theme CSS -->
    <link href="css/freelancer.css" rel="stylesheet">

    <!-- Custom Fonts -->
    <link href="vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css?family=Montserrat:400,700" rel="stylesheet" type="text/css">
    <link href="https://fonts.googleapis.com/css?family=Lato:400,700,400italic,700italic" rel="stylesheet" type="text/css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <!-- Adding jQuery and jQuery file uploader -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.min.js"></script>
    <!-- For uploading images using jQuery -->
    <script type="text/javascript" src="./libs/jQFU/js/vendor/jquery.ui.widget.js"></script>
    <script type="text/javascript" src="./libs/jQFU/js/jquery.iframe-transport.js"></script>
    <script type="text/javascript" src="./libs/jQFU/js/jquery.fileupload.js"></script>
    <script type="text/javascript" src="./js/ajaxify.js"></script>
    <script type="text/javascript" src="./js/serialize.js"></script>
    <script type="text/javascript" src="./js/fileupload.js"></script>

</head>

<body id="page-top" class="index">

    <!-- This is used to upload the Files via jQuery, so don't modify plz -->
    <div style="display: none">
        <form>
            <input type="hidden" name="no_html" value="yes">
            <input type="hidden" name="type" value="mod">
            <input type="hidden" name="username" id="modUsername">
            <input id="fileupload" type="file" name="files[]" accept="*/*">
        </form>
        <form>
            <input type="hidden" name="no_html" value="yes">
            <input type="hidden" name="type" value="tex">
            <input type="hidden" name="username"  id="texUsername">
            <input id="textureupload" type="file" name="files[]" accept="image/*">
        </form>
    </div>

    <!-- Navigation -->
    <nav id="mainNav" class="navbar navbar-default navbar-fixed-top navbar-custom">
        <div class="container">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header page-scroll">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                    <span class="sr-only">Toggle navigation</span> Menu <i class="fa fa-bars"></i>
                </button>
                <a class="navbar-brand" href="#page-top">Nebula</a>
            </div>

            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav navbar-right">
                    <li class="hidden">
                        <a href="#page-top"></a>
                    </li>

                    <li class="page-scroll">
                        <a href="#portfolio">Create a listing</a>
                    </li>

                   <!-- <li class="page-scroll">
                        <a href="#portfolio">Product Templates</a>
                    </li>-->

                    <li class="page-scroll">
                        <a href="#about">About</a>
                    </li>
                    <li class="page-scroll">
                        <a href="#contact">Contact</a>
                    </li>
                    
                    <li>
                        <a data-toggle="modal" style="cursor: pointer;" data-target="#createModal12" id="createbutton">Edit a listing</a>
                    </li>
                    <li>
                        <a data-toggle="modal" style="cursor: pointer;" data-target="#createModal13" id="deletebutton">Delete a listing</a>
                    </li>
                    <li>
                        <a href="/logout.php">Log out</a>
                    </li>

                </ul>
            </div>
            <!-- /.navbar-collapse -->
        </div>
        <!-- /.container-fluid -->
    </nav>

    <!-- Header -->
    <header>
        <div class="container">
            <div class="row">
                <div class="col-lg-12">
                    <img class="img-responsive" src="img/nebula_logo_v2.png" alt="">
                    <div class="intro-text">
                        <!--<span class="name">Nebula</span> -->
                        <hr class="star-light">
                        <span class="skills">An Augmented Reality Storefront</span>
                    </div>
                </div>
            </div>
        </div>
    </header>

    <!-- Portfolio Grid Section -->
    <section id="portfolio">
        <div class="container">
            <div class="row">
                <div class="col-lg-12 text-center">
                    <h2>Create Listing</h2>

                    <hr class="star-primary">

                    <button data-toggle="modal" data-target="#createModal" id="createbutton" type="button" class="btn btn-primary btn-lg">Create Custom Listing</button>

                   
                &nbsp
                &nbsp

                <br />


                  <!--  <h2> OR <h2>
                   
                    <h3>Select a category</h2>

                </div>
            </div>
            
            <div class="row">
                <div class="col-sm-4 portfolio-item">
                    <a href="#Furniture" class="portfolio-link" data-toggle="modal">
                        <div class="caption">
                            <div class="caption-content">
                                <i class="fa fa-search-plus fa-3x"></i>
                            </div>
                        </div>
                        <img src="img/portfolio/couch.jpg" class="img-responsive" alt="">
                    </a>
                </div>


                <div class="col-sm-4 portfolio-item">
                    <a href="#books" class="portfolio-link" data-toggle="modal">
                        <div class="caption">
                            <div class="caption-content">
                                <i class="fa fa-search-plus fa-3x"></i>
                            </div>
                        </div>
                        <img src="img/portfolio/book.jpg" class="img-responsive" alt="">
                    </a>
                </div>

            
        

            <div class="row">
                <div class="col-sm-4 portfolio-item">
                    <a href="#Music" class="portfolio-link" data-toggle="modal">
                        <div class="caption">
                            <div class="caption-content">
                                <i class="fa fa-search-plus fa-3x"></i>
                            </div>
                        </div>
                        <img src="img/portfolio/instruments.jpg" class="img-responsive" alt="">
                    </a>
                </div>
            </div>
        </div>-->

    </section>

    <!-- About Section -->
    <section class="success" id="about">
        <div class="container">
            <div class="row">
                <div class="col-lg-12 text-center">
                    <h2>About</h2>
                    <hr class="star-light">
                </div>
            </div>
            <div class="row">
                <div class="col-lg-4 col-lg-offset-2">
                    <p>Nebula is an Augmented Reality platform giving Hololens users the ability to shop for various products. With the Hololens platform it gives users the opportunity to virtually touch, move, and position life-sized products around their current environment.</p>
                </div>
                <div class="col-lg-4">
                    <p>This application and web store was .</p>
                </div>
                <div class="col-lg-8 col-lg-offset-2 text-center">
 
                </div>
            </div>
        </div>
    </section>

    <!-- Contact Section -->
    <section id="contact">
        <div class="container">
            <div class="row">
                <div class="col-lg-12 text-center">
                    <h2>Contact Us</h2>
                    <hr class="star-primary">
                </div>
            </div>
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2">
                    <!-- To configure the contact form email address, go to mail/contact_me.php and update the email address in the PHP file on line 19. -->
                    <!-- The form should work on most web servers, but if the form is not working you may need to configure your web server differently. -->
                    <form name="sentMessage" id="contactForm" novalidate>
                        <div class="row control-group">
                            <div class="form-group col-xs-12 floating-label-form-group controls">
                                <label>Name</label>
                                <input type="text" class="form-control" placeholder="Name" id="name" required data-validation-required-message="Please enter your name.">
                                <p class="help-block text-danger"></p>
                            </div>
                        </div>
                        <div class="row control-group">
                            <div class="form-group col-xs-12 floating-label-form-group controls">
                                <label>Email Address</label>
                                <input type="email" class="form-control" placeholder="Email Address" id="email" required data-validation-required-message="Please enter your email address.">
                                <p class="help-block text-danger"></p>
                            </div>
                        </div>
                        <div class="row control-group">
                            <div class="form-group col-xs-12 floating-label-form-group controls">
                                <label>Phone Number</label>
                                <input type="tel" class="form-control" placeholder="Phone Number" id="phone" required data-validation-required-message="Please enter your phone number.">
                                <p class="help-block text-danger"></p>
                            </div>
                        </div>
                        <div class="row control-group">
                            <div class="form-group col-xs-12 floating-label-form-group controls">
                                <label>Message</label>
                                <textarea rows="5" class="form-control" placeholder="Message" id="message" required data-validation-required-message="Please enter a message."></textarea>
                                <p class="help-block text-danger"></p>
                            </div>
                        </div>
                        <br>
                        <div id="success"></div>
                        <div class="row">
                            <div class="form-group col-xs-12">
                                <button type="submit" class="btn btn-success btn-lg">Send</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </section>

    <!-- Footer -->
    <footer class="text-center">
        <div class="footer-above">
            <div class="container">
                <div class="row">
                    <div class="footer-col col-md-4">
                        <h3>Location</h3>
                        <p>3141 Chesnut Street
                            <br>Philadelphia, PA 19104</p>
                    </div>
                    <div class="footer-col col-md-4">
                        <h3>Around the Web</h3>
                        <ul class="list-inline">
                            <li>
                                <a href="#" class="btn-social btn-outline"><i class="fa fa-fw fa-facebook"></i></a>
                            </li>
                            <li>
                                <a href="#" class="btn-social btn-outline"><i class="fa fa-fw fa-google-plus"></i></a>
                            </li>
                            <li>
                                <a href="#" class="btn-social btn-outline"><i class="fa fa-fw fa-twitter"></i></a>
                            </li>
                            <li>
                                <a href="#" class="btn-social btn-outline"><i class="fa fa-fw fa-linkedin"></i></a>
                            </li>
                            <li>
                                <a href="#" class="btn-social btn-outline"><i class="fa fa-fw fa-dribbble"></i></a>
                            </li>
                        </ul>
                    </div>
                    <div class="footer-col col-md-4">
                        <h3>About Nebula</h3>
                        <p>Nebula is a free to use, open source application . <a href="http://startbootstrap.com">Insert Team Website</a>.</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="footer-below">
            <div class="container">
                <div class="row">
                    <div class="col-lg-12">
                        Copyright &copy; Your Website 2016
                    </div>
                </div>
            </div>
        </div>
    </footer>

    <!-- Scroll to Top Button (Only visible on small and extra-small screen sizes) -->
    <div class="scroll-top page-scroll hidden-sm hidden-xs hidden-lg hidden-md">
        <a class="btn btn-primary" href="#page-top">
            <i class="fa fa-chevron-up"></i>
        </a>
    </div>


    

    <!-- jQuery -->
    <!--<script src="vendor/jquery/jquery.min.js"></script>-->

    <!-- Bootstrap Core JavaScript -->
    <script src="vendor/bootstrap/js/bootstrap.min.js"></script>

    <!-- Plugin JavaScript -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-easing/1.3/jquery.easing.min.js"></script>

    <!-- Contact Form JavaScript -->
    <script src="js/jqBootstrapValidation.js"></script>
    <script src="js/contact_me.js"></script>

    <!-- Theme JavaScript -->
    <script src="js/freelancer.min.js"></script>

    <script src="script.js"></script>

</body>

</html>

<div id="createModal" class="modal fade" role ="dialog">
                <div class="modal-dialog modal-lg">
                    <!--Modal content-->
                    <div class="modal-content">
                        <div class="modal-header" style="text-align: center">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Create New Listing</h4>
                        </div>
                        <div class="modal-body">
                            <div style="text-align: left">
                                
                                <!--
                                <form enctype='application/json' style="text-align: center" method="post" name="form">
                                -->
                                <form name="submitForm">
                                    <input name="seller" value="<?= $user['username'] ?>" type="text" class="form-control" placeholder="Username" disabled> &nbsp
                                    <input name="email" value="<?= $user['email'] ?>" type="text" class="form-control" placeholder="Email" disabled> &nbsp
                                    <select name="category" value="" class="form-control">
                                        <option selected disabled value="choose">--Category--</option>
                                        <option value="furniture">Furniture</option>
                                        <option value="books">Books</option>
                                        <option value="music">Music</option>
                                    </select> &nbsp
                                    <input name="item" value="" type="text" class="form-control" placeholder="Item Name"> &nbsp
                                    <input name="itemdesc" value="" type="text" class="form-control" placeholder="Item Description"> &nbsp
                                    <input name="model" value="" type="hidden" class="form-control" placeholder="Model"> &nbsp
                                    <input name="texture" value="" type="hidden" class="form-control" placeholder="Texture"> &nbsp
                                    <input name="price" value="" type="text" class="form-control" placeholder="Price ($00.00)"> &nbsp
                                </form>

                                    <h2> Image Upload </h2>
                                    <!--<form enctype = "multipart/form-data" action="uploaded.php" method="post">
                                        Only JPEG, PNG, JPG types allowed. .
                                        <br/><br/>
                                        <!--<div id="filediv"><input name="file[]" type="file" id="file"/></div>-->
                                        <br/>

                                        <div class="form-group">
                                            <input type="hidden" name="filename" id="imageName">
                                            <input type="hidden" name="texturename" id="textureName">
                                            <div class="col-md-10">
                                                <span class="btn btn-success fileinput-button image-button">
                                                    <i class="glyphicon glyphicon-plus"></i>
                                                    <span>Select Model File...</span>
                                                </span>
                                                <span id="modelName"></span>
                                                <br>
                                                <br>
                                                <!-- The global progress bar -->
                                                <div id="progress" class="progress" style="width: 465px">
                                                    <div class="progress-bar progress-bar-success"></div>
                                                </div>
                                                <!-- The container for the uploaded files -->
                                                <div id="files" class="files"></div>
                                            </div>

                                            <div class="col-md-10">
                                                <span class="btn btn-success fileinput-button image-buttonx">
                                                    <i class="glyphicon glyphicon-plus"></i>
                                                    <span>Select Texture File...</span>
                                                </span>
                                                <span id="texName"></span>

                                                <br>
                                                <br>
                                                <!-- The global progress bar -->
                                                <div id="progressx" class="progress" style="width: 465px">
                                                    <div class="progress-bar progress-bar-success"></div>
                                                </div>
                                                <!-- The container for the uploaded files -->
                                                <div id="texturefiles" class="files"></div>
                                            </div>
                                        </div>


                            <!-- <input type="button" id="add_more" class="upload" value="Add More Files"/>-->
                            <!--<input type="submit" value="Upload File" name="uploadimage" id="upload" class="upload"/>-->
                        <!--</form>-->

                        <style>
                            @import "http://fonts.googleapis.com/css?family=Droid+Sans";
                            form{
                            background-color:#fff
                            }
                            #maindiv{
                            width:960px;
                            margin:10px auto;
                            padding:10px;
                            font-family:'Droid Sans',sans-serif
                            }
                            #formdiv{
                            width:500px;
                            float:left;
                            text-align:center
                            }
                            form{
                            padding:40px 20px;
                            box-shadow:0 0 10px;
                            border-radius:2px
                            }
                            h2{
                            margin-left:30px
                            }
                            .upload{
                            background-color:blue;
                            border:1px solid blue;
                            color:#fff;
                            border-radius:5px;
                            padding:10px;
                          
                            }
                            .upload:hover{
                            cursor:pointer;
                            background:black;
                            border:1px solid #c20b0b;
                            box-shadow:0 0 5px rgba(0,0,0,.75)
                            }
                            #file{
                            color:green;
                            padding:5px;
                            border:1px dashed #123456;
                            background-color:#f9ffe5
                            }
                            #upload{
                            margin-left:45px
                            }
                            #noerror{
                            color:green;
                            text-align:left
                            }
                            #error{
                            color:red;
                            text-align:left
                            }
                            #img{
                            width:17px;
                            border:none;
                            height:17px;
                            margin-left:-20px;
                            margin-bottom:91px
                            }
                            .abcd{
                            text-align:center
                            }
                            .abcd img{
                            height:100px;
                            width:100px;
                            padding:5px;
                            border:1px solid #e8debd
                            }
                            b{
                            color:red
                            }
                            </style>

                        <!--- Including PHP Script -->
                        <?php /*include "upload.php";*/ ?>


                                <div class="modal-footer">
                                   <!-- <button class="btn btn-default" id="sendAjax">Submit</button> -->
                                </div>
                                </form>
                    </div>
              <button class="btn btn-default" id="sendAjax">Submit</button>
                </div>
            </div>
        </div>
</div>


<div id="createModala" class="modal fade" role ="dialog">
                <div class="modal-dialog modal-lg">
                    <!--Modal content-->
                    <div class="modal-content">
                        <div class="modal-header" style="text-align: center">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Create New account</h4>
                        </div>
                        <div class="modal-body">
                            <div style="text-align: left">

                                <form enctype='application/json' style="text-align: center" method="post" name="form">
                                    <input name="name" value="" type="text" class="form-control" placeholder="First and Last Name"> &nbsp
                                    <input name="email" value="" type="text" class="form-control" placeholder="Email"> &nbsp
                                    <input name="pw" value="" type="password" class="form-control" placeholder="Password"> &nbsp
                                    <input name="pw2" value="" type="password" class="form-control" placeholder="Confirm password"> &nbsp
                                    &nbsp
                                    

                              <!--  <form id="upload" action="upload.php" method="POST" enctype="multipart/form-data">
                                    <fieldset>

                                        <input type="hidden" id="MAX_FILE_SIZE" name="MAX_FILE_SIZE" value="300000" />

                                        <div>
                                            <label for="fileselect">Files to upload:</label>
                                            <input type="file" id="fileselect" name"fileselect[]" multiple="multiple" />
                                            <div id="filedrag">or drop files here</div>
                                        </div>

                                        <div id="submitbutton">
                                            <button type="submit">Upload Files</button>
                                    </fieldset>
                                </form>

                                <div id="messages">
                                    <p>Status</p>
                                </div>

                                <style>
                                #filedrag
                                {
                                    display: none;
                                    font-weight: bold;
                                    text-align: center;
                                    padding: 1em 0;
                                    margin: 1em 0;
                                    color: #555;
                                    border: 2px dashed #555;
                                    border-radius: 7px;
                                    cursor: default;
                                }

                                #filedrag.hover
                                {
                                    color: #f00;
                                    border-color: #f00;
                                    border-style: solid;
                                    box-shadow: inset 0 3px 4px #888;
                                }

                                </style>-->

                                <div class="modal-footer">
                                    <button class="btn btn-default">Create</button>
                                </div>
                                </form>
                    </div>
                </div>
            </div>
        </div>
</div>

<!--<div id="createModalF" class="modal fade" role ="dialog">
                <div class="modal-dialog modal-lg">
                    <!--Modal content-
                    <div class="modal-content">
                        <div class="modal-header" style="text-align: center">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Create New Furniture Listing</h4>
                        </div>
                        <div class="modal-body">
                            <div style="text-align: left">

                                <form enctype='application/json' style="text-align: center" method="post" name="form">
                                    <input name="firstname" value="" type="text" class="form-control" placeholder="First Name"> &nbsp
                                    <input name="lastname" value="" type="text" class="form-control" placeholder="Last Name"> &nbsp
                                    <input name="email" value="" type="text" class="form-control" placeholder="Email"> &nbsp
                                    <select name="category" value="" class="form-control">
                                        <option selected disabled value="choose">Furniture</option>
                                        <option value="furniture">Furniture</option>
                                        <option value="books">Books</option>
                                        <option value="music">Music</option>
                                    </select> &nbsp
                                    <input name="item" value="" type="text" class="form-control" placeholder="Item Name"> &nbsp
                                    <input name="itemdesc" value="" type="text" class="form-control" placeholder="Item Description"> &nbsp
                                    <input name="price" value="" type="text" class="form-control" placeholder="Price ($00.00)"> &nbsp

                                <div class="modal-footer">
                                    <button name="submit" onclick ="onsubmit()" class="btn btn-default">Submit</button>
                                </div>
                                </form>
                    </div>
                </div>
            </div>
        </div>
</div>

<div id="createModalB" class="modal fade" role ="dialog">
                <div class="modal-dialog modal-lg">
                    <!--Modal content--
                    <div class="modal-content">
                        <div class="modal-header" style="text-align: center">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Create New Book Listing</h4>
                        </div>
                        <div class="modal-body">
                            <div style="text-align: left">

                                <form enctype='application/json' style="text-align: center" method="post" name="form">
                                    <input name="firstname" value="" type="text" class="form-control" placeholder="First Name"> &nbsp
                                    <input name="lastname" value="" type="text" class="form-control" placeholder="Last Name"> &nbsp
                                    <input name="email" value="" type="text" class="form-control" placeholder="Email"> &nbsp
                                    <select name="category" value="" class="form-control">
                                        <option selected disabled value="choose">Books</option>
                                        <option value="furniture">Furniture</option>
                                        <option value="books">Books</option>
                                        <option value="music">Music</option>
                                    </select> &nbsp
                                    <input name="item" value="" type="text" class="form-control" placeholder="Item Name"> &nbsp
                                    <input name="itemdesc" value="" type="text" class="form-control" placeholder="Item Description"> &nbsp
                                    <input name="price" value="" type="text" class="form-control" placeholder="Price ($00.00)"> &nbsp

                                <div class="modal-footer">
                                    <button name="submit" onclick ="onsubmit()" class="btn btn-default">Submit</button>
                                </div>
                                </form>
                    </div>
                </div>
            </div>
        </div>
</div>

<div id="createModalM" class="modal fade" role ="dialog">
                <div class="modal-dialog modal-lg">
                    <!--Modal content--
                    <div class="modal-content">
                        <div class="modal-header" style="text-align: center">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Create New Music Listing</h4>
                        </div>
                        <div class="modal-body">
                            <div style="text-align: left">

                                <form enctype='application/json' style="text-align: center" method="post" name="form">
                                    <input name="firstname" value="" type="text" class="form-control" placeholder="First Name"> &nbsp
                                    <input name="lastname" value="" type="text" class="form-control" placeholder="Last Name"> &nbsp
                                    <input name="email" value="" type="text" class="form-control" placeholder="Email"> &nbsp
                                    <select name="category" value="" class="form-control">
                                        <option selected disabled value="choose">Music</option>
                                        <option value="furniture">Furniture</option>
                                        <option value="books">Books</option>
                                        <option value="music">Music</option>
                                    </select> &nbsp
                                    <input name="item" value="" type="text" class="form-control" placeholder="Item Name"> &nbsp
                                    <input name="itemdesc" value="" type="text" class="form-control" placeholder="Item Description"> &nbsp
                                    <input name="price" value="" type="text" class="form-control" placeholder="Price ($00.00)"> &nbsp

                                <div class="modal-footer">
                                    <button  onclick ="onsubmit()" class="btn btn-default">Submit</button>
                                </div>
                                </form>
                    </div>
                </div>
            </div>
        </div>
</div> -->

<script src="js/json.js"></script>
<!--<script src="js/filedrag.js"></script>-->

                <!-- <script>
                                var form = document.forms['form'];

                                form.onsubmit = function (e) {
                                    //stop regular form submission
                                    e.preventDefault();

                                    //collect the form data 
                                    var data = {};
                                    for (var i = 0, ii = form.length; i <ii; ++i) {
                                        var input = form[i];
                                        if (input.name) {
                                            data[input.name] = input.value;
                                        }
                                    }

                                    //construct an HTTP request
                                    var xhr = new XMLHttpRequest();
                                    xhr.open(form.method, form.action, true);
                                    xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');

                                    //send the collected data as JSON
                                    xhr.send(JSON.stringify(data));

                                    console.log(JSON.stringify(data));

                                    xhr.onloadend = function () {
                                        //done
                                    };
                                };

                </script>-->
