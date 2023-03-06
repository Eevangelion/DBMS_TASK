import './App.css';
import Feed from "./pages/Feed";
import UserPage from "./pages/UserPage";
import SearchPage from "./pages/SearchPage";
import Subscribes from './pages/Subscribes';
import Settings from './components/Settings/Settings';
import TagRedactor from './components/TagRedactor/TagRedactor';
import { Route, Routes, useLocation } from 'react-router-dom';
import CreateJoke from "./components/CreateJoke/CreateJoke"
import CreateReport from './components/CreateReport/CreateReport';
import ErrorPage from './pages/ErrorPage';
import Subscribe from './components/Subscribe/Subscribe';
import ReportsList from './components/ReportsList/ReportsList';
import AuthPage from './pages/Auth';
import OAuthRedirect from './components/OAuthRedirect/OAuthRedirect';
import JokeModal from './components/JokeModal/JokeModal';
import RegisterPage from './pages/Register';
import BannedPage from './pages/BannedPage';
import Unsubscribe from './components/Unsubscribe/Unsubscribe';
import {AuthProvider, RequireAuth, RequireAdmin, RequireNotBanned} from './context/context';
import DevelopSettings from './components/DevelopSettings/DevelopSettings';
const App = () => {
  let location = useLocation();
  let state = location.state;
  return (
    <AuthProvider>
      <Routes location={state?.backgroundLocation || location}>
        <Route index element={<RequireAuth>
                                <RequireNotBanned>
                                  <Feed />
                                </RequireNotBanned>
                              </RequireAuth>}/>
        <Route path='feed/' element={<RequireAuth>
                                      <RequireNotBanned>
                                        <Feed />
                                      </RequireNotBanned>
                                    </RequireAuth>}/>
        <Route path='subscribes/' element={<RequireAuth>
                                            <RequireNotBanned>
                                              <Subscribes />
                                            </RequireNotBanned>
                                          </RequireAuth>}/>
        <Route path='user/:username' element={<RequireAuth>
                                                <RequireNotBanned>
                                                  <UserPage />
                                                </RequireNotBanned>
                                              </RequireAuth>}/>
        <Route path='search/:type/:query?' element={<RequireAuth>
                                                      <RequireNotBanned>
                                                        <SearchPage />
                                                      </RequireNotBanned>
                                                    </RequireAuth>}/>
        <Route path='login/' element={<AuthPage />} />
        <Route path='register/' element={<RegisterPage />} />
        <Route path='oauth/' element={<OAuthRedirect /> } />
        <Route path="tagredactor/" element={<RequireAdmin>
                                              <RequireNotBanned>
                                                <TagRedactor />
                                              </RequireNotBanned>
                                            </RequireAdmin>} />
        <Route path="reportslist/" element={<RequireAdmin>
                                              <RequireNotBanned>
                                                <ReportsList />
                                              </RequireNotBanned>
                                            </RequireAdmin>}/>
        <Route path="banned/" element={<RequireAuth>
                                          <BannedPage />
                                      </RequireAuth>}/>
        <Route path="*" element={<ErrorPage />}/>
      </Routes>
      {state?.backgroundLocation && (
        <Routes>
            <Route path="/create_joke" element={<RequireAuth>
                                                  <RequireNotBanned>
                                                    <CreateJoke />
                                                  </RequireNotBanned>
                                                </RequireAuth>} />
            <Route path="/settings" element={<RequireAuth>
                                              <RequireNotBanned>
                                                <Settings />
                                              </RequireNotBanned>
                                            </RequireAuth>} />
            <Route path="/develop_settings" element={<RequireAdmin>
                                                      <RequireNotBanned>
                                                        <DevelopSettings/>
                                                      </RequireNotBanned>
                                                    </RequireAdmin>} />
            <Route path="/create_report/:jokeID" element={<RequireAuth>
                                                            <RequireNotBanned>
                                                              <CreateReport />
                                                            </RequireNotBanned>
                                                          </RequireAuth>} />
            <Route path="/subscribe/:receiverID" element={<RequireAuth>
                                                            <RequireNotBanned>
                                                              <Subscribe />
                                                            </RequireNotBanned>
                                                          </RequireAuth>}/>
            <Route path="/unsubscribe/:receiverID" element={<RequireAuth>
                                                              <RequireNotBanned>
                                                                <Unsubscribe />
                                                              </RequireNotBanned>
                                                            </RequireAuth>} />
            <Route path="/joke/:jokeID" element={<RequireAuth>
                                                  <RequireNotBanned>
                                                    <JokeModal />
                                                  </RequireNotBanned>
                                                </RequireAuth>} />
        </Routes>
      )}
    </AuthProvider>
  )
}

export default App;
