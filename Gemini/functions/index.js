const functions = require("firebase-functions");
const admin = require("firebase-admin");
// const nodemailer = require("nodemailer");
// const axios = require("axios");
admin.initializeApp();

exports.getUserInfo = functions.https.onRequest(async (req, res) => {
  const userId = req.query.userId;

  try {
    // Firebase Authenticationでユーザー情報を取得
    const userRecord = await admin.auth().getUser(userId);
    const userInfo = {
      uid: userRecord.uid,
      email: userRecord.email,
      displayName: userRecord.displayName,
      photoURL: userRecord.photoURL,
    };
    res.status(200).json(userInfo);
  } catch (error) {
    console.error("Error fetching user data:", error);
    res.status(500).send("Error fetching user data");
  }
});

exports.sayHello = functions.https.onRequest((req, res) => {
  const name = req.query.name || "World";
  res.send(`Hello, ${name}!`);
});

// すべてのユーザーのメールアドレスを取得
exports.getAllEmails = functions.https.onRequest(async (req, res) => {
  const usersInfo = []; // メールアドレスを格納する配列
  try {
  // Firebase Authenticationのすべてのユーザーを取得
    const listAllUsers = async (nextPageToken) => {
      const result = await admin.auth()
          .listUsers(1000, nextPageToken); // 1ページあたり1000件まで取得
      result.users.forEach((userRecord) => {
        if (userRecord.email) {
          usersInfo.push({
            uid: userRecord.uid, // UID
            email: userRecord.email, // メールアドレス
          });
        }
      });
      // 次のページがある場合は再帰的に処理
      if (result.pageToken) {
        await listAllUsers(result.pageToken);
      }
    };
    // 最初のページから開始
    await listAllUsers();
    res.status(200).json(usersInfo); // メールアドレスの配列を返す
  } catch (error) {
    console.error("Error fetching user emails:", error);
    res.status(500).send("Error fetching user emails");
  }
});

// メール送信用の設定
// const transporter = nodemailer.createTransport({
//   service: "gmail",
//   auth: {
//     user: "7310wave@gmail.com",
//     pass: "jdcs pjwj xndr buec",
//   },
// });

// 全ての UID を取得して予測データを送信
exports.sendPredictsByEmail = functions.https.onRequest(async () => {
  try {
  // Firebase 管理者でユーザーリストを取得
    const users = await admin.auth().listUsers();
    for (const userRecord of users.users) {
      const uid = userRecord.uid;
      const email = userRecord.email;
      if (!email) {
        console.log(`UID ${uid} のメールアドレスが見つかりません。`);
        continue;
      }
      console.log(`UID ${uid} のメールアドレスは ${email} です。`);
      // UID に基づいてカスタムトークンを生成
    //   const token = await admin.auth().createCustomToken(uid);
    //   // Go サーバーの API 呼び出し
    //   const response = await axios.get(
    //       "https://backend-71857953091.us-central1.run.app/api/predict",
    //       {
    //         headers: {
    //           "Authorization": `Bearer ${token}`, // トークンをヘッダーに含める
    //           "Content-Type": "application/json",
    //         },
    //       },
    //   );
    //   const {productIds} = response.data;
    //   // メールを送信
    //   const mailOptions = {
    //     from: "7310wave@gmail.com",
    //     to: email,
    //     subject: "予測結果",
    //     text: `UID: ${uid}\n予測された商品ID: ${productIds.join(", ")}`,
    //   };
    //   await transporter.sendMail(mailOptions);
    //   console.log(`メールを ${email} に送信しました。`);
    }
  } catch (error) {
    console.error("エラーが発生しました:", error);
  }
});
