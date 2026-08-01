import { apiFetch } from "./client";

export async function getTaskList() {
  return apiFetch("/api/v1/tasks", {
    method: "GET",
  });
}

export async function createTask(tasks) {
  return apiFetch("/api/v1/tasks", {
    method: "POST",
    body: JSON.stringify(tasks),
  });
}

export async function getTask(id) {
  return apiFetch(`/api/v1/tasks/${id}`, {
    method: "GET",
  });
}

export async function updateTask(id, tasks) {
  return apiFetch(`/api/v1/tasks/${id}`, {
    method: "PUT",
    body: JSON.stringify(tasks),
  });
}

export async function deleteTask(id) {
  return apiFetch(`/api/v1/tasks/${id}`, {
    method: "DELETE",
  });
}

export async function patchTaskComplete(id) {
  return apiFetch(`/api/v1/tasks/${id}/complete`, {
    method: "PATCH",
  });
}
