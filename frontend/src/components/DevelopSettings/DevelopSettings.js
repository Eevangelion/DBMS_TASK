import { useNavigate, useLocation, Link } from "react-router-dom"
import { useGetUserByIDQuery } from "../../services/Joke";
import './DevelopSettings.css';

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


const DevelopSettings = () => {
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
            Доступ к данному ресурсу ограничен.
        </div>
    );

    const developSettingsWindow = (
        <div className="modal-window">
            <div className="buttons">
                <Link to={`/tagredactor`} 
                      style={linkStyle}>
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

export default DevelopSettings;