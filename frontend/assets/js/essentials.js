k = Cookies.get("key")
// (k)
if(k != null && k != "" && k != undefined){
}else{
   window.location.replace('/');
}

$(".logout").click(function(){
    Cookies.remove("key")
    window.location.replace('/');
})