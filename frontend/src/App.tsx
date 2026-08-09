import { Route, Routes } from "react-router-dom";
import { LoginPage } from "./LoginPage";

function App() {
  return (
    <Routes>
      <Route path="login" element={<LoginPage></LoginPage>}></Route>
      <Route path="/" element={<div>部屋ビュー（仮）</div>} />
      <Route path="/list" element={<div>タスク一覧（仮）</div>} />
      <Route path="/focus" element={<div>フォーカスモード（仮）</div>} />
    </Routes>
  );
}

export default App;
