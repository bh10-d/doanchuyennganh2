import React, { useState, useEffect } from "react";
import { Todo } from "./Todo";
import { TodoForm } from "./TodoForm";
// import { v4 as uuidv4 } from "uuid";
import { EditTodoForm } from "./EditTodoForm";

export const TodoWrapper = () => {
  const [todos, setTodos] = useState([]);
  const [success, isSuccess] = useState('');
  
  const BASE_URL='http://localhost:8080/v1/items'

  useEffect(() =>{
    fetch(BASE_URL).then(res => res.json()).then(data => {
      setTodos(data.data);
      // console.log(data.data[0].title)
      // console.log(data)
    });
  },[success])

  const addTodo = async (data) => {
    await fetch(BASE_URL, {
      method: 'POST',
      body: JSON.stringify({title: data}),
      headers: {
        'Content-Type': 'application/json'
      }
    });
    // const result = await response.json();
    // console.log(result);
    isSuccess(data)
    // console.log(data);
  }
  
  const deleteTodo = async (id) => {
    console.log(id)
    fetch(`${BASE_URL}/${id}`, {
      method: 'DELETE'
    })
    .then(res => res.json())
    .then(data => {
      // console.log(data)
      isSuccess(data)
    })
    .catch(error => console.error(error));
    
  };

  const toggleComplete = (id) => {
    setTodos(
      todos.map((todo) =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  }

  const editTodo = (id) => {
    setTodos(
      todos.map((todo) =>
        todo.id === id ? { ...todo, isEditing: !todo.isEditing } : todo
      )
    );
  }

  const editTask = async (task, id) => {
    // setTodos(
    //   todos.map((todo) =>
    //     todo.id === id ? { ...todo, task, isEditing: !todo.isEditing } : todo
    //   )
    // );
    await fetch(`${BASE_URL}/${id}`, {
      method: 'PUT',
      body: JSON.stringify({title: task}),
      headers: {
        'Content-Type': 'application/json'
      }
    });

    isSuccess(task)
    // console.log(task, id)
  };

  return (
    <div className="TodoWrapper">
      <h1>An Active Day !</h1>
      <TodoForm addTodo={addTodo} />
      {/* display todos */}
      {(todos.length !== 0 || todos !== undefined)?
        todos.map((todo) =>
        todo.isEditing ? (
          <EditTodoForm editTodo={editTask} task={todo} />
        ) : (
          <Todo
            key={todo.id}
            task={todo}
            deleteTodo={deleteTodo}
            editTodo={editTodo}
            toggleComplete={toggleComplete}
          />
        )
      ):<></>
      }
    </div>
  );
};
