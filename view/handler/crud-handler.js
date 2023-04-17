const form = document.getElementById('upload-form');

form.addEventListener('submit', function(event) {
    event.preventDefault();

    const xhr = new XMLHttpRequest();
    xhr.open('POST', 'http://localhost:8080/files');

    xhr.onload = function() {
        console.log(xhr.responseText);
    };

    xhr.send(new FormData(form));
});