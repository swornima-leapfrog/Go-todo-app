async function deleteTodo(event) {
  try {
    const id = event.currentTarget.getAttribute("id");

    await axios.delete(`http://localhost:3000/${id}`);

    window.location.reload();
  } catch (error) {
    console.log(error);
  }
}

async function addTodo() {
  try {
    const todo = {
      title: document.getElementById("add-todo").value
    };

    if (!todo.title) {
      alert("Please enter a todo");
      return;
    }

    await axios.post("http://localhost:3000", todo);

    window.location.reload();
  } catch (error) {
    console.log(error);
  }
}

async function toggleTodo(event) {
  try {
    const id = event.currentTarget.getAttribute("id");

    const todo = {
      done: event.currentTarget.checked
    };

    await axios.put(`http://localhost:3000/${id}`, todo);

    window.location.reload();
  } catch (error) {
    console.log(error);
  }
}

async function markAllAsRead() {
  try {
    await axios.put("http://localhost:3000/all");

    window.location.reload();
  } catch (error) {
    console.log(error);
  }
}
