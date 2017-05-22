function updateUsername(){
    jQuery('#modUsername').val(jQuery('input[name=seller]').val());
    jQuery('#texUsername').val(jQuery('input[name=seller]').val());
}

jQuery(document).ready(function() {

    
        //Assign the username to the suitable fields and make sure if something changes to update that
        jQuery('#modUsername').val(jQuery('input[name=seller]').val());
        jQuery('#texUsername').val(jQuery('input[name=seller]').val());
        jQuery('input[name=seller]').on("change", updateUsername);

        //This is triggered when we click on the submit button for create listing.
        jQuery("#sendAjax").click(function(e){
            if (
                !$('select[name="category"]').val() ||
                !$('input[name="item"]').val() ||
                !$('input[name="itemdesc"]').val() ||
                !$('input[name="price"]').val() ||
                !$('input[name="texturename"]').val() ||
                !$('input[name="filename"]').val()
                ) {
                    alert('Please complete all fields');
                    e.PreventDefault();
                    return false;
            }

            var form = document.forms['submitForm'];
            //e.preventDefault();

            
            var API_URL = 'https://thingproxy.freeboard.io/fetch/http://40.71.214.175:5073/additem';

            //Update model filename
            jQuery('input[name=model]').val(jQuery("#imageName").val());
            //Update the texture filename
            jQuery('input[name=texture]').val(jQuery("#textureName").val());
            
            //collect the form data
            var data = {};
            for (var i = 0, ii = form.length; i <ii; ++i) {
                var input = form[i];
                if (input.name) {
                    data[input.name] = input.value;
                }
            }

            

            jQuery.ajax({
                url: API_URL,
                type: 'POST',
                crossDomain: true,
                data: JSON.stringify(data),
                dataType: 'text',
                contentType: "application/json",
                success: function (response) {
                    //var resp = JSON.parse(response);
                    //console.log(resp.statusText);
                    alert('Post Successful');
                },
                error: function (/*xhr,*/ status) {
                    //console.log(status);

                    alert("Post failed for some reason");
                }
            });

            $('#createNewListing')[0].reset();
            $('#createModal').modal('hide');
            $('#progress, #progressx').css('width', '0px');
            $('#texName').html('');
            $('#modelName').html('');
        });

        //File Upload for the model
        jQuery('.image-button').click(function(){
            jQuery('#fileupload').click();
        });

        //File Upload for the texture
        jQuery('.image-buttonx').click(function(){
            jQuery('#textureupload').click();
        });

        
        
        var url = './upload_image.php';
        

        jQuery('#fileupload').fileupload({
            url: url,
            dataType: 'json',
            maxNumberOfFiles: 1,
            done: function (e, data) {
                jQuery.each(data.result.files, function (index, file) {

                    //Save the uploaded image file name to be used for later
                    image_src   = 'http://www.kieranplante.com/uploads/'+file.name;
                    original    = image_src.replace('_mod_'+jQuery('input[name=seller]').val(), '');
                    


                    jQuery("#imageName").val(image_src);
                    jQuery('#editModel').val(image_src);
                    jQuery("#modelName").html(original);

                    jQuery('#viewModel').attr('href', image_src);
                    jQuery('#viewModel').html(original);
                   

                });
            },
            progressall: function (e, data) {
                var progress = parseInt(data.loaded / data.total * 100, 10);
                jQuery('#progress .progress-bar').css(
                    'width',
                    progress + '%'
                );
            }
        }).prop('disabled', !jQuery.support.fileInput)
            .parent().addClass(jQuery.support.fileInput ? undefined : 'disabled');

        
        var url = './upload_texture.php';
        

        jQuery('#textureupload').fileupload({
            url: url,
            dataType: 'json',
            maxNumberOfFiles: 1,
            done: function (e, data) {
                jQuery.each(data.result.files, function (index, file) {

                    //Save the uploaded image file name to be used for later
                    image_src = 'http://www.kieranplante.com/uploads/'+file.name;
                    original    = image_src.replace('_tex_'+jQuery('input[name=seller]').val(), '');

                    jQuery("#textureName").val(image_src);
                    jQuery('#editTexture').val(image_src);

                    jQuery("#texName").html(original);

                    jQuery('#viewTexture').html(original);
                    jQuery('#viewTexture').attr('href', image_src);
                });
            },
            progressall: function (e, data) {
                var progress = parseInt(data.loaded / data.total * 100, 10);
                jQuery('#progressx .progress-bar').css(
                    'width',
                    progress + '%'
                );
            }
        }).prop('disabled', !jQuery.support.fileInput)
            .parent().addClass(jQuery.support.fileInput ? undefined : 'disabled');
    });
