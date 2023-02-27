import { Link, useLocation } from "react-router-dom";
import { useSelector } from "react-redux";
import {useGetUserByIDQuery, useGetJokeByIDQuery, useApplyReportMutation, useRemoveReportMutation } from "../../services/service";
import "./Report.css";

const Report = (props) => {
    const reportsListIsActive = useSelector(state => state.pagesReducer.reportListIsActive);
    const location = useLocation();
    const {
        data: receiver,
        isLoading: loadingReceiver,
    } = useGetUserByIDQuery(props.report.receiver_id);
    const {
        data: sender,
        isLoading: loadingSender,
    } = useGetUserByIDQuery(props.report.sender_id);
    const {
        data: joke,
        isLoading: loadingJoke,
    } = useGetJokeByIDQuery(props.report.receiver_joke_id);
    const [applyReport] = useApplyReportMutation();
    const [denyReport] = useRemoveReportMutation();
    if (loadingReceiver || loadingSender || loadingJoke) {
        return (
            <div classname="report-post">
                Загрузка...
            </div>
        )
    }

    return (
        <div className="report-post" style={reportsListIsActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <div className="report-top-panel" style={reportsListIsActive ? {} : {borderBottom: "0.1vh solid #555"}}>
                <div className="report-sender">
                    От кого: <Link to={`/user/${sender.name}`} style={reportsListIsActive ? {} : {color: "#050585"}}>{sender.name}</Link>
                </div>
                <div className="report-receiver"> 
                    Кому: <Link to={`/user/${receiver.name}`} style={reportsListIsActive ? {} : {color: "#050585"}}>{receiver.name}</Link>
                </div>
                <Link   to={`/joke/${joke.id}`}
                        state={{ backgroundLocation: location }}
                        className="joke-link"
                        style={reportsListIsActive ? {} : {color: "#050585"}}
                        onClick={(event) => {if (!reportsListIsActive) event.preventDefault()}}
                >
                    Просмотреть шутку
                </Link>
            </div>
            <div className="report-description">
                Описание: {props.report.description}
            </div> 
            <div className="buttons">
                <Link 
                        className="apply-button"
                        style={reportsListIsActive ? {} : {backgroundColor: "#118", color: "#aaa"}}
                        onClick={(event)=>{
                            if (reportsListIsActive) applyReport(props.report.id)
                            else event.preventDefault();
                        }}
                >Одобрить</Link>
                <Link 
                        className="deny-button"
                        style={reportsListIsActive ? {} : {backgroundColor: "#666", color: "#aaa"}}
                        onClick={(event)=>{
                            if (reportsListIsActive) denyReport(props.report.id)
                            else event.preventDefault();
                        }}
                >Отклонить</Link>
            </div>
        </div>
    )
}

export default Report;