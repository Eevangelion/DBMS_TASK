import './App.css';
import Main from "./pages/Main";
import UserPage from "./pages/UserPage";
import { Route, Routes } from 'react-router-dom';

const App = () => {
  return (
    <Routes>
      <Route index element={<Main />}/>
      <Route path='user/:username' element={<UserPage />}/>
    </Routes>
  )
}

export default App;
