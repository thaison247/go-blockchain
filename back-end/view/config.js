$(document).ready(function () {
  $("#blockList").empty();

  let req = callApi(
    "http://localhost:4000/generateMnemonicPhrase",
    "GET",
    null
  );
  req.then(
    (rs) => {
      console.log("dfasdfas");
      console.log({ rs });
    },
    (err) => {
      console.log("err: ", err);
    }
  );

  var socket = io("http://localhost:4000");
  socket.on("connect", function () {});
});

function callApi(urlReq, method, dataBody) {
  return new Promise((resolve, reject) => {
    var request = $.ajax({
      url: urlReq,
      method: method,
      data: JSON.stringify(dataBody),
      contentType: "application/json",
    });

    request.done(function (msg) {
      return resolve(msg.data);
    });

    request.fail(function (jqXHR, textStatus) {
      return reject(jqXHR.responseJSON.data);
    });
  });
}
