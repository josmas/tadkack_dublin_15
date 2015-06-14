document.addEventListener('DOMContentLoaded', function() {
  if (Notification.permission !== "granted")
    Notification.requestPermission();
});
$(function() {
  var mysip = ""
  $.phono({
    apiKey: "fbf2e6a4279b8b8d242057c17378d8af",
    audio: {
      type: "flash"
    },
    onReady: function(event) {
      mysip = this.sessionId;
      $('#sip').text("sip:" + this.sessionId);
      var url = "/register?sip=" + encodeURIComponent(this.sessionId)
      $.get("/register?sip=" + url, function(data) {
        console.log("Registered! " + data);
        $('#status').text("Wait for call");
      });
    },
    phone: {
      onIncomingCall: function(event) {
        var call = event.call;
        $('#status').text("incoming call!");
        $('#btnanswer').removeClass('hidden');
        var notification = new Notification('Notification title', {
          icon: 'https://cdn0.iconfinder.com/data/icons/basic-ui-elements-round/700/08_phone-512.png',
          body: "Incomming call! " + call.id,
        });
        notification.onclick = function() {
          $('#btnanswer').addClass('hidden');
          $('#btnhangup').removeClass('hidden');
          call.answer();
        };
        $("#btnanswer").click(function() {
          $('#btnhangup').removeClass('hidden');
          $('#btnanswer').addClass('hidden');
          call.answer();
        });
        $("#btnhangup").click(function() {
          call.hangup();
          $('#btnanswer').removeClass('hidden');
          $('#btnhangup').removeClass('hidden');
          $('#status').text("Wait for call");
        });
        call.bind({
          onHangup: function(event) {
            $('#btnanswer').addClass('hidden');
            $('#btnhangup').addClass('hidden');
            $('#status').text("Wait for call");
            console.log("Call hung up");
          }
        });
      }
    }
  });
});
