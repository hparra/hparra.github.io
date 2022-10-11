# Auditing an AWS Deployment

This is organized by AWS Console areas e.g. IAM, Billing & Cost Management, etc.

## Prequisites

- AWS Console access
- [AWS CLI](https://aws.amazon.com/cli/)

## TODO

### Identity and Access Management (IAM)

First **review [best practices for IAM](http://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)**. Keep these in mind when auditing. Make lists about what needs to be checked, and what needs to be changed, but DO NOT make changes until you are absolutely sure there are no unwanted side-effects. It's easy to create downtime with the click of button!

Review of IAM Best Practices:
- Lock away your AWS account (root) access keys
- Create individual IAM users
- Use groups to assign permissions to IAM users
- Grant least privilege
- Configure a strong password policy for your users
- Enable MFA for privileged users
- Use roles for applications that run on Amazon EC2 instances
- Delegate by using roles instead of by sharing credentials
- Rotate credentials regularly
- Remove unnecessary credentials
- Use policy conditions for extra security
- Monitor activity in your AWS account

Remember that a user does not necessarily need to be a person. Applications can be users, but also consider using Roles instead.

**Create and analyze a [Credential Report](https://console.aws.amazon.com/iam/home#credential_report)**:
- Check if `<root_account>` access keys are being used
- Check all access key usage for latest date and service used
- Check for stale accounts - is the application no longer in use? Does the employee no longer work there?
- Check for accounts whose use is not obvious - ask around

If multiple users and/or applications are using the same access keys then neither the Access Advisor nor the Credential Report provide sufficient information. Check if [CloudTrail](https://console.aws.amazon.com/cloudtrail/home) is available. This may not be necessary if you are auditing code repositories as well, but for larger projects this may be the only comprehensive solution.

**Make a list of users and their access keys.** This is useful when analyzing application deployments and server environments (and code repositories - *gasp*) for credentials used.

### Billing & Cost Management

**Check the [Cost Explorer](https://console.aws.amazon.com/billing/home#/costexplorer)**. Enable it if necessary - it does not cost anything.
