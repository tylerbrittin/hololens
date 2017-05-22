<?php

require_once './libs/UploadHandler.php';

define('UPLOAD_DIR',	getcwd().'/uploads/');
define('UPLOAD_PATH',	'/uploads/');

$files      = new UploadHandler(array(	'upload_dir'	=> UPLOAD_DIR,
    					'upload_url'	=> UPLOAD_PATH,
));

$filename           = $files->name;
die($filename);

