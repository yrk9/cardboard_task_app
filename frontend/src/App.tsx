import { Route, Routes } from "react-router-dom";
import { LoginPage } from "./LoginPage";
import { TaskListPage } from "./TaskListPage";

function App() {
  return (
    <Routes>
      <Route path="login" element={<LoginPage></LoginPage>} />
      <Route path="/" element={<TaskListPage></TaskListPage>} />
      <Route path="/view" element={<div>部屋ビュー（仮）</div>} />
      <Route path="/focus" element={<div>フォーカスモード（仮）</div>} />
    </Routes>
  );
}

export default App;
