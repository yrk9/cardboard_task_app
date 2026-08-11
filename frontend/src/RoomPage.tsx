import React, { useState, useEffect } from "react";
import { useAuth } from "./context/AuthContext";
import { getTaskList } from "./api/tasks";
import { useNavigate } from "react-router-dom";
import { Task } from "./types";

function getRoomStage(count: number): { emoji: string; message: string } {
  if (count === 0) return { emoji: "✨", message: "最高の状態です！" };
  if (count <= 3) return { emoji: "📦", message: "まだ余裕あり" };
  if (count <= 7) return { emoji: "📦📦", message: "そろそろ片付けましょう" };
  if (count <= 12) return { emoji: "🏠", message: "足の踏み場が…" };
  return { emoji: "😱", message: "助けてください" };
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
      <p>{stage.emoji}</p>
      <p>{stage.message}</p>
      <p>未完了タスク:{incompleteCount}件</p>
    </div>
  );
}
