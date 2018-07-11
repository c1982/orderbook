  function getTIMESTAMP(timestamp) {
    var date = new Date(timestamp * 1000);
    var year = date.getFullYear();
    var month = ("0"+(date.getMonth()+1)).substr(-2);
    var day = ("0"+date.getDate()).substr(-2);
    var hour = ("0"+date.getHours()).substr(-2);
    var minutes = ("0"+date.getMinutes()).substr(-2);
    var seconds = ("0"+date.getSeconds()).substr(-2);

    return year+"-"+month+"-"+day+" "+hour+":"+minutes+":"+seconds;
  }

function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
};