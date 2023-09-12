// This is what we will send and how we will send the data;
class Event {
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

const sendMessage = (eventName, payload) => {
  const event = new Event(eventName, payload);
  connection.send(JSON.stringify(event));
};

// This is how we handle the events we get fromt the server;
const routeHandler = (event) => {
  if (!event.type) {
    console.log("Undefined Event Type;");
    return;
  }
  switch (event.type) {
    case "new_message":
      console.log("New Message");

    default:
      console.log("Unknown event tytpe;");
  }
};

const changeChatRoom = () => {
  chatroom = document.getElementById("chatroom").value;
  //   To make the form not to refresh when submit; ( Not to navigate to the other URL )
  return false;
};

const sendMessages = () => {
  chatroomMessage = document.getElementById("chatroomMessage").value;
  if (chatroomMessage) {
    // We are sending an event witht the type of message_sent
    sendMessage("message_sent", chatroomMessage);
  }
  return false;
};

window.onload = () => {
  const changeChatRoomForm = document.getElementById("changeChatroomForm");
  const chatroomMessagesForm = document.getElementById("chatroomMessagesForm");

  // Add events to the above refrenced elements;
  changeChatRoomForm.onsubmit = changeChatRoom;
  chatroomMessagesForm.onsubmit = sendMessages;

  //    Sockets
  if (window["WebSocket"]) {
    console.log("Browser is capable of using the websockets");
    connection = new WebSocket(`ws://${document.location.host}/ws`);
    connection.onmessage = (event) => {
      alert(event.data);
    };
    console.log(document.location.host);
  } else {
    alert("Your version of the browser doesnt suport the websockets;");
  }
};
