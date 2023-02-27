import { useSelector } from "react-redux";
import { useGetReportsQuery } from "../../services/service";
import TopPanel from "../TopPanel/TopPanel";
import Report from "../Report/Report";
import styles from "./ReportList.module.css";


const ReportsList = () => {
    const reportsListIsActive = useSelector(state => state.pagesReducer.reportListIsActive);
    const {
        data: response,
        isLoading: loadingReports
    } = useGetReportsQuery();
    if (loadingReports) {
        return <div>Загрузка...</div>;
    }
    const {reports, amount} = response;

    const reportPosts = reports.map((report) => 
    {
        return <Report report={report} />
    });
    if (!reports) {
        return <div className={styles.mainPage}>
                    <TopPanel />
                    <div className={styles.info} style={reportsListIsActive ? {} : {backgroundColor: "#676a6c"}}>
                        <div className={styles.feed}>
                            <div className={styles.txt}>Жалоб нет</div>
                        </div>
                        {/* <PageSelector pageState={true} /> */}
                    </div>
                </div>;
    }
    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={reportsListIsActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    <div className={styles.reportsList}>
                        {reportPosts}
                    </div>
                </div>
                {/* <PageSelector pageState={true} /> */}
            </div>
        </div>
    )
}

export default ReportsList;