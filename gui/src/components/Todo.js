import React from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons'
import { faTrash, faCheck } from '@fortawesome/free-solid-svg-icons'

export const Todo = ({task, deleteTodo, editTodo, toggleComplete}) => {
  return (
    <div className="Todo">
        {/* <p className={`${task.status ? "completed" : "incompleted"}`} onClick={() => toggleComplete(task.id)}>{task.task}</p> */}
        <p className={task.id} onClick={() => toggleComplete(task.id)}>{task.title}</p>
        <div>
        <FontAwesomeIcon className="edit-icon" icon={faPenToSquare} onClick={() => editTodo(task.id)} />
        <FontAwesomeIcon className="delete-icon" icon={faCheck} onClick={() => deleteTodo(task.id)} />
        </div>
    </div>
  )
}
