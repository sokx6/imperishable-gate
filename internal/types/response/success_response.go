package response

// 常见成功响应
var PongResponse = Response{Message: "pong"}
var RegisterSuccessResponse = Response{Message: "Registered successfully"}
var AddLinkSuccessResponse = Response{Message: "Added successfully"}
var AddNamesSuccessResponse = Response{Message: "Names added successfully"}
var AddRemarkByLinkSuccessResponse = Response{Message: "Remark added successfully"}
var AddRemarkByNameSuccessResponse = Response{Message: "Remark added successfully"}
var AddTagsByLinkSuccessResponse = Response{Message: "Tags added successfully"}
var AddTagsByNameSuccessResponse = Response{Message: "Tags added successfully"}
var DeleteSuccessResponse = Response{Message: "Links deleted successfully"}
var DeleteTagsByNameSuccessResponse = Response{Message: "Tags deleted successfully"}
var DeleteNamesByLinkSuccessResponse = Response{Message: "Names deleted successfully"}
var DeleteTagsByLinkSuccessResponse = Response{Message: "Tags deleted successfully"}
var WatchSuccessResponse = Response{Message: "Link is now being watched"}
var UnwatchSuccessResponse = Response{Message: "Link is no longer being watched"}
var EmailVerifiedSuccessResponse = Response{Message: "Email verified successfully!"}
var VerificationEmailResentSuccessResponse = Response{Message: "Verification email resent successfully!"}
var RegistrationSuccessResponse = Response{Message: "Registration successful. Please check your email to verify your account."}
var PasswordResetSuccessResponse = Response{Message: "Password has been reset successfully. You can now log in with your new password."}
var SendResetPasswordEmailSuccessResponse = Response{Message: "Reset password email sent successfully."}
var ResetPasswordEmailSentSuccessResponse = Response{Message: "If the email is registered, a reset password email has been sent."}
