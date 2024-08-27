import { useState, useEffect } from "react";
import TaskInputComponent from "./components/taskinput"
import TaskComponent from "./components/task"
import { getTodos, addTodo, updateTodo, deleteTodo } from "./api/todoApi"

export default function App() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    async function fetchTasks() {
      try {
        let todos = await getTodos()
        setTasks(todos)
      } catch (error) {
        console.error("Error fetching tasks:", error)
      }
    }
    fetchTasks()
  }, [])

  async function addTask(text) {
    try {
      await addTodo({ "id": 0, "text": text })
      let todos = await getTodos()
      setTasks(todos)
    } catch (error) {
      console.error('Error adding todo:', error);
    }
  }

  async function updateTask(task) {
    try {
      await updateTodo(task)
      let todos = await getTodos()
      setTasks(todos)
    } catch (error) {
      console.error('Error updating todo:', error);
    }
  };

  async function removeTask(taskId) {
    try {
      await deleteTodo(taskId)
      let todos = await getTodos()
      setTasks(todos)
    } catch (error) {
      console.error('Error deleting todo:', error);
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-gradient-to-r from-cyan-500 to-blue-500">
      <div className="flex flex-col gap-2 items-center w-2/6 bg-white p-6 rounded">
        <h2 className="text-xl">TODO App:</h2>
        <TaskInputComponent onUpdate={addTask}/>
        {tasks.map(task => (
          <TaskComponent key={task.id} task={task} onUpdate={updateTask} onComplete={removeTask} />
        ))}
      </div>
    </div>
  );
}