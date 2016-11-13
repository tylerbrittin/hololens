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
                        /*var xhr = new XMLHttpRequest();
                        xhr.open(form.method, url, true);
                        xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');

                        //send the collected data as JSON
                        xhr.send(JSON.stringify(data));

                        console.log(JSON.stringify(data));

                        xhr.onloadend = function () {*/

                        


                            $.ajax({
                                url: url,
                                type: 'POST',
                                data: JSON.stringify(data),
                                contentType: "application/json",
                                dataType: 'jsonp',
                            });

                        //console.log(JSON.stringify(data))
                        //console.log(url);

                };

