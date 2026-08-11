import { Route, Routes } from "react-router-dom";
import { LoginPage } from "./LoginPage";
import { TaskListPage } from "./TaskListPage";
import { RoomPage } from "./RoomPage";

function App() {
  return (
    <Routes>
      <Route path="login" element={<LoginPage></LoginPage>} />
      <Route path="/" element={<TaskListPage></TaskListPage>} />
      <Route path="/view" element={<RoomPage></RoomPage>} />
      <Route path="/focus" element={<div>フォーカスモード（仮）</div>} />
    </Routes>
  );
}

export default App;
