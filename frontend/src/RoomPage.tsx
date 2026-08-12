import React, { useState, useEffect } from "react";
import { useAuth } from "./context/AuthContext";
import { getTaskList } from "./api/tasks";
import { useNavigate } from "react-router-dom";
import { Task } from "./types";
import "./RoomPage.css";

function getRoomStage(count: number): { emoji: string; message: string } {
  if (count === 0) return { emoji: "✨", message: "最高の状態です！" };
  if (count <= 3) return { emoji: "📦", message: "まだ余裕あり" };
  if (count <= 7) return { emoji: "📦📦", message: "そろそろ片付けましょう" };
  if (count <= 12) return { emoji: "🏠", message: "足の踏み場が…" };
  return { emoji: "😱", message: "助けてください" };
}

function getBoxClassName(task: Task): string {
  const sizeClass =
    task.priority === "high"
      ? "box-high"
      : task.priority === "mid"
        ? "box-mid"
        : "box-low";
  return `box ${sizeClass}`.trim();
}

export function RoomPage() {
  const { token, isLoading } = useAuth();
  const navigate = useNavigate();
  const [tasks, setTasks] = useState<Task[]>([]);
  const incompleteCount = tasks.filter(
    (task) => task.completed_at === null,
  ).length;
  const stage = getRoomStage(incompleteCount);

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

  return (
    <div>
      <h1>ダン部屋</h1>
      <div className="room">
        {tasks.map((task) => (
          <div
            key={task.id}
            className={getBoxClassName(task)}
            title={task.title}
          />
        ))}
      </div>

      <p>{stage.message}</p>
      <p>未完了タスク:{incompleteCount}件</p>
    </div>
  );
}
