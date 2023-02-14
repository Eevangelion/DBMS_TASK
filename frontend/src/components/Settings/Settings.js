import { useNavigate, useLocation, Link } from "react-router-dom"
import { useGetUserByIDQuery } from "../../services/Joke";
import './Settings.css';

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "380px",
    height: "30px",
    borderRadius: "45px",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
}


const Settings = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const userID = localStorage.getItem('userID');

    const {
        data: user,
        isLoading: loadingUser, 
    }= useGetUserByIDQuery(userID);

    if (loadingUser) {
        return <div className="modal-window">Загрузка...</div>;
    }

    const settingsWindow = (
        <div className="modal-window">
            <div className="buttons">
                <button className="back-button" onClick={() => navigate(-1)}>
                    Назад
                </button>
            </div>
        </div>
    );

    const developSettingsWindow = (
        <div className="modal-window">
            <div className="buttons">
                <Link to={`/developsettings`} 
                      style={linkStyle}
                      state={{ backgroundLocation: location }}>
                    <strong>Настройки разработчика</strong>
                </Link>
                <button className="back-button" onClick={() => navigate(-1)}>
                    Назад
                </button>
            </div>
        </div>
    );

    if (user.role === "admin") {
        return developSettingsWindow;
    } else {
        return settingsWindow;
    }
}

export default Settings;