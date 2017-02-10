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
                        var url = 'http://40.87.66.169:5073/additem'

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

