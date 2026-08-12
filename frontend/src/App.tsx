import { Route, Routes, useNavigate } from "react-router-dom";
import { LoginPage } from "./LoginPage";
import { TaskListPage } from "./TaskListPage";
import { RoomPage } from "./RoomPage";

function App() {
  const navigate = useNavigate();

  const handleTransfer = (buttonType: string) => {
    if (buttonType == "login") {
      navigate("/login");
    } else if (buttonType == "tasklist") {
      navigate("/");
    } else if (buttonType == "view") {
      navigate("/view");
    } else if (buttonType == "focus") {
      navigate("/focus");
    } else {
      navigate("/");
    }
  };

  return (
    <Routes>
      <Route path="login" element={<LoginPage></LoginPage>} />
      <Route
        path="/"
        element={<TaskListPage handleTransfer={handleTransfer}></TaskListPage>}
      />
      <Route
        path="/view"
        element={<RoomPage handleTransfer={handleTransfer}></RoomPage>}
      />
      <Route path="/focus" element={<div>フォーカスモード（仮）</div>} />
    </Routes>
  );
}

export default App;
