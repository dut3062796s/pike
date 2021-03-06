import React from "react";
import Configs from "../configs";
import axios from "axios";
import { message, Switch } from "antd";

import { getAlarmI18n, getCommonI18n } from "../../i18n";
import { ALARMS_TRY } from "../../urls";

const category = "alarms";

const columns = [
  {
    title: getAlarmI18n("name"),
    dataIndex: "name"
  },
  {
    title: getAlarmI18n("uri"),
    dataIndex: "uri"
  },
  {
    title: getAlarmI18n("template"),
    dataIndex: "template",
    width: 400,
    render: row => {
      if (!row) {
        return;
      }
      return (
        <pre
          style={{
            "white-space": "pre-wrap",
            "word-wrap": "break-word"
          }}
        >
          {row}
        </pre>
      );
    }
  },
  {
    title: getCommonI18n("enabled"),
    dataIndex: "enabled",
    render: row => {
      return <Switch disabled={true} defaultChecked={row} />;
    }
  },
  {
    title: getCommonI18n("description"),
    dataIndex: "description"
  }
];

const fields = [
  {
    label: getAlarmI18n("name"),
    key: "name",
    placeholder: getAlarmI18n("namePlaceHolder"),
    rules: [
      {
        required: true,
        message: getAlarmI18n("nameRequireMessage")
      }
    ],
    options: ["upstream", "uncaught-error"],
    type: "select"
  },
  {
    label: getAlarmI18n("uri"),
    key: "uri",
    placeholder: getAlarmI18n("uriPlaceHolder"),
    rules: [
      {
        required: true,
        message: getAlarmI18n("uriRequireMessage")
      }
    ]
  },
  {
    label: getAlarmI18n("template"),
    key: "template",
    placeholder: getAlarmI18n("templatePlaceHolder"),
    rules: [
      {
        required: true,
        message: getAlarmI18n("templateRequireMessage")
      }
    ],
    type: "textarea"
  },
  {
    label: getCommonI18n("enabled"),
    key: "enabled",
    type: "switch"
  },
  {
    label: getCommonI18n("description"),
    key: "description",
    type: "textarea",
    placeholder: getCommonI18n("descriptionPlaceholder")
  }
];

class Alarms extends Configs {
  constructor(props) {
    super(props);
    Object.assign(this.state, {
      title: getAlarmI18n("createUpdateTitle"),
      description: getAlarmI18n("createUpdateDescription"),
      category,
      columns,
      fields,
      actions: [
        {
          key: "test",
          text: getAlarmI18n("try"),
          icon: "thunderbolt",
          handle: async row => {
            try {
              await new Promise(resolve => setTimeout(resolve, 5000));
              await axios.post(ALARMS_TRY.replace(":name", row.name));
              message.info("done");
            } catch (err) {
              message.error(err.message);
            }
          }
        }
      ]
    });
  }
}

export default Alarms;
