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

           //console.log(JSON.stringify(data));
                        //construct an HTTP request
                        var url = //this URL has been taken out

                            $.ajax({
                                url: url,
                                type: 'POST',
                                crossDomain: true,
                                data: JSON.stringify(data),
                                dataType: 'json',
                                contentType: "application/json",
                                success: function (response) {
                                    var resp = JSON.parse(response)
                                    alert(resp.status);
                                },
                                error: function (xhr, status) {
                                    alert("error");
                                }
                            });
                        //console.log(JSON.stringify(data))
                        //console.log(url);


                };

