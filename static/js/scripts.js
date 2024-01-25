document.addEventListener("htmx:configRequest", function (event) {
  // Disable the button before making the htmx request

  const button = document.getElementById("myButton");
  button.classList.add("disabled-button");
  console.log("button disabled");
});
