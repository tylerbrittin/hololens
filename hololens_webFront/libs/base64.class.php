<?php
/**
 * Author: Aboubakr Seddik Ouahabi (aboubakr[at]codernix.com)
 */
 

define('UPLOAD_DIR',	getcwd().'/uploads/');

class base64{

    static public function exportBase64($filename){

        //Get the absolute path of the image
        $fileDestination    = UPLOAD_DIR.$filename;

        //Let's convert it to base64
        $fileData = base64_encode(file_get_contents($fileDestination));
        //Making sure, this will be recognized as an image
        $data = 'data: '.mime_content_type($fileDestination).';charset=utf8;base64,'.$fileData;

        return($data);
    }
}

?>
