<?php

if (isset($_POST['submit'])) {
	$j = 0; 		// Variable for indexing uploaded image
	$target_path = "uploads/";  // Declaring path for uploaded images.
	for ( $i = 0; $i < count($_FILES['file']['name']); $i++) {
		// Loop to get individual element from the array

		$validextensions = array("jpeg","jpg","png");  // Extensions that are allowed
		$ext = explode('.', basename($_FILES['file']['name'][$i]))  // Explode file name from dot (.)
		$file_extension = end($ext); // Store extentions in the variable
		$target_path = $target_path . md5(uniqid()) . "." . $ext[count($ext) - 1]; //Set the target path with the name of the image
		$j = $j + 1; // Increment the number of uploaded images according to the files in the array
		if (($_FILES["file"]["size"][$i] < 100000) // Approx. 100kb files can be uploaded
			&& in_array($file_ectensions, $validextensions)) {

			if (move_uploaded_file($_FILES['file']['tmp_name'][$i], $target_path)) {

				// If file moved to uploads folder
				echo $j. ').<span id="noerror">Image uploaded successfully!. </span><br/><br/>';
			} else {	// If File was not moved.
				echo $j. '). <span id="error">Please try again!.<span><br/><br/>';
			}
		} else {		// If file size and File type incorrect.
			echo $j. ').<span id="error"> Invalid file Size or Type</span><br/><br/>';
		}
	}
}
?>