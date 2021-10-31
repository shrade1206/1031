function send(){
fetch('http://127.0.0.1:8080/inser', {
method: 'POST',
body: JSON.stringify({title : document.getElementById("set").value})
})
.then((res)=> res.json())
.then(res => {
document.getElementById('list').innerText = res
})

}