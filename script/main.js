let nama = document.querySelector("#nama")
let usia = document.querySelector("#usia")
let btn = document.querySelector("#btn")
let hasil = document.querySelector("#hasil")


btn.addEventListener("click", function() {
    if(nama.value === "") {
        alert("data wajib di isi")
    } else {
        hasil.textContent = nama.value;
    }

})