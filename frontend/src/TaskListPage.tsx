import React, { useState, useEffect } from "react";
import { useAuth } from "./context/AuthContext";
import {
  createTask,
  deleteTask,
  getTaskList,
  patchTaskComplete,
} from "./api/tasks";
import { useNavigate } from "react-router-dom";
import { Task } from "./types";

export function TaskListPage() {
  const { token, isLoading } = useAuth();
  const navigate = useNavigate();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [errorMessage, setErrorMessage] = useState<string | null>(null); //エラーメッセージ表示用
  //   const [isLoadingTasks, setIsLoadingTask] = useState<boolean>(false);
  const [form, setForm] = useState({
    title: "",
    description: "",
    priority: "mid",
    due_date: "",
  });

  async function fetchTasks() {
    const data = await getTaskList();
    setTasks(data ?? []);
  }

  useEffect(() => {
    //ローディングでない状態でトークンがない場合(ログインできてないので弾く)
    if (!isLoading && !token) {
      navigate("/login");
    }
    fetchTasks();
  }, [isLoading, token, navigate]);

  function handleChange(
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
    console.log(form);
  }

  async function handleAddTask(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setErrorMessage(null);

    try {
      await createTask({
        ...form,
        due_date: form.due_date ? `${form.due_date}T00:00:00Z` : null,
      });
      await fetchTasks();
    } catch (error) {
      setErrorMessage("タスクを追加できませんでした");
    }
  }

  async function handleDeleteTask(id: number) {
    setErrorMessage(null);

    try {
      await deleteTask(id);
      await fetchTasks();
    } catch (error) {
      setErrorMessage("タスクを削除できませんでした");
    }
  }

  async function handleCompleteTask(id: number) {
    setErrorMessage(null);

    try {
      await patchTaskComplete(id);
      await fetchTasks();
    } catch (error) {
      setErrorMessage("タスクを完了できませんでした");
    }
  }

  return (
    <div>
      <h1>ダン部屋</h1>
      <form onSubmit={handleAddTask}>
        <label>
          タイトル
          <input
            name="title"
            value={form.title}
            onChange={handleChange}
          ></input>
        </label>
        <label>
          内容
          <input
            name="description"
            value={form.description}
            onChange={handleChange}
          ></input>
        </label>
        <label>
          優先度
          <select name="priority" value={form.priority} onChange={handleChange}>
            <option value="high">高</option>
            <option value="mid">中</option>
            <option value="low">低</option>
          </select>
        </label>
        <label>
          締め切り日
          <input
            type="date"
            name="due_date"
            value={form.due_date}
            onChange={handleChange}
          ></input>
        </label>

        <button>タスクの追加</button>
        {errorMessage && <p>{errorMessage}</p>}
      </form>
      <div>
        {tasks != null &&
          tasks.map((task: Task) => (
            <li key={task.id}>
              {task.title},{task.description}, {task.priority}, {task.due_date},{" "}
              {task.completed_at != null ? task.completed_at : "未完了"}
              <button onClick={() => handleDeleteTask(task.id)}>
                タスクの削除
              </button>
              <button onClick={() => handleCompleteTask(task.id)}>完了</button>
            </li>
          ))}
      </div>
    </div>
  );
}
