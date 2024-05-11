import React, { useEffect, useMemo, useState } from "react";
import styles from "./activity.module.css";
import { TypeOfActivities } from "../../data/types";

type ActivityProps = {
	title: TypeOfActivities;
	activities: Array<TypeOfActivities>;
	setActivities: (activities: Array<TypeOfActivities>) => void;
};

export const Activity = (props: ActivityProps) => {
	const { title, activities, setActivities } = props;
	const [active, setActive] = useState(false);

	let classesForBlock = styles.wrapper__activity;
	let classesForText = styles.activity;
	let indexAct = -1;

	activities.forEach((act, index) => {
		if (act.description === title.description) {
			indexAct = index;
		}
	});

	useEffect(() => {
		if (indexAct !== -1) {
			setActive(true);
		}
	}, []);

	if (indexAct !== -1) {
		classesForBlock += " " + styles.wrapper__activity__active;
		classesForText += " " + styles.activity__active;
	}

	useEffect(() => {
		if (active) {
			classesForBlock += " " + styles.wrapper__activity__active;
			classesForText += " " + styles.activity__active;
			if (indexAct === -1) {
				setActivities([...activities, title]);
			}
		} else {
			classesForBlock = styles.wrapper__activity;
			classesForText = styles.activity;
			if (indexAct !== -1) {
				setActivities(activities.filter((elem) => elem !== title));
			}
		}
	}, [active]);

	return (
		<div
			className={classesForBlock}
			key={title.description}
			onClick={(event) => {
				event.stopPropagation();
				setActive(!active);
			}}
		>
			<h4 className={classesForText}>{title.description}</h4>
		</div>
	);
};
