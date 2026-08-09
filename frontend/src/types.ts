export type Task = {
  id: number;
  title: string;
  descriptions: string;
  priority: "high" | "mid" | "low";
  due_date: string | null;
  completed_at: string | null;
  created_at: string;
  updated_at: string;
};
