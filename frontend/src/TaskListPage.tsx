import { useState, useEffect } from "react";
import { useAuth } from "./context/AuthContext";
import { createTask, getTaskList } from "./api/tasks";
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
    descriptions: "",
    priority: "mid",
    due_date: "",
  });

  useEffect(() => {
    async function fetchTasks() {
      const data = await getTaskList();
      setTasks(data);
    }
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
  }

  async function handleAddTask(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setErrorMessage(null);

    try {
      await createTask({ ...form, due_date: form.due_date || null });
      const data = await getTaskList();
      setTasks(data);
    } catch (error) {
      setErrorMessage("タスクを追加できませんでした");
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
            name="descriptions"
            value={form.descriptions}
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
        {tasks.map((task: Task) => (
          <li key={task.id}>
            {task.title},{task.descriptions}, {task.priority}, {task.due_date}
          </li>
        ))}
      </div>
    </div>
  );
}
