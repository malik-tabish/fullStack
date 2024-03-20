const baseUrl = "http://127.0.0.1:8080/";
window.onload = () => {
  const form = document.getElementById("form");
  form.addEventListener("submit", function (event) {
    event.preventDefault();
    const myFormData = new FormData(form);
    const formDataJSON = {};
    myFormData.forEach((value, key) => {
      formDataJSON[key] = value;
    });
    postRequest(formDataJSON);
  });

  displayData();

  async function deleteUser(userId) {
    try {
      const response = await fetch(baseUrl + "users/" + userId, {
        method: "DELETE",
      });
      if (response.ok) {
        console.log("User deleted successfully");
        // Optionally, update the UI to remove the deleted row
      } else {
        console.error("Failed to delete user:", response.statusText);
      }
    } catch (error) {
      console.error("There was a problem with your fetch operation:", error);
    }
    displayData();
  }
};

async function displayData() {
  console.log("hi");
  try {
    const res = document.getElementById("tbody");
    const response = await fetch(baseUrl + "users");
    const message = await response.json();
    message.forEach((element) => {
      res.innerHTML += `
          <tr>
            <td>${element.id}</td>
            <td>${element.name}</td>
            <td>${element.username}</td>
            <td>${element.password}</td>
            <td>
              <i class="bi bi-pencil-square pointer" onclick="edit(${element.id})"></i>
              <i class="bi bi-trash pointer" onclick="deleteUser(${element.id})"></i>
            </td>
          </tr>
        `;
    });
  } catch (error) {
    console.error("There was a problem with your fetch operation:", error);
  }
}

async function deleteUser(userId) {
  try {
    const response = await fetch(baseUrl + "users/" + userId, {
      method: "DELETE",
    });
    if (response.ok) {
      console.log("User deleted successfully");
      displayData();
      // Optionally, update the UI to remove the deleted row
    } else {
      console.error("Failed to delete user:", response.statusText);
    }
  } catch (error) {
    console.error("There was a problem with your fetch operation:", error);
  }
}
