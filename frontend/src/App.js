import './App.css';
import Feed from "./pages/Feed";
import UserPage from "./pages/UserPage";
import SearchPage from "./pages/SearchPage";
import { Route, Routes, useLocation } from 'react-router-dom';
import CreateJoke from "./components/CreateJoke/CreateJoke"

const App = () => {

  let location = useLocation();
  let state = location.state;
  return (
    <>
      <Routes location={state?.backgroundLocation || location}>
        <Route index element={<Feed />}/>
        <Route path='feed' element={<Feed />}/>
        <Route path='user/:username' element={<UserPage />}/>
        <Route path='search/?:q&:t' element={<SearchPage />}/>
      </Routes>
      {state?.backgroundLocation && (
        <Routes>
            <Route path="/create_joke" element={<CreateJoke />} />
        </Routes>
      )}
    </>
  )
}

export default App;
