//删除信息
function logout() {
    // Create a new XMLHttpRequest object
    var xhr = new XMLHttpRequest();
    // Define the request URL and method
    var url = 'http://101.42.251.239:8080/userLogout';
    var method = 'POST';
    // Set up the request
    xhr.open(method, url, true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    // Create the request body
    var requestBody = JSON.stringify({});
    // Send the request
    xhr.send(requestBody);
    // Set up the callback function to handle the response
    xhr.onload = function () {
        if (xhr.status === 200) {
            var response = JSON.parse(xhr.responseText);
            return response.ret
        }else {
            return -1
        }
    };
    window.location.href = "http://101.42.251.239:8080"
};