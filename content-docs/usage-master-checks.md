# Checks Tab

This page allows you to create and manage alerts.

Each check go through a list of hosts and make decision based on check expressions.


# How to create a new check?

Click the `New Check` button. It will ask you to fill in a form.

* **Name:** The name of your check.

* **Interval:** How frequently the master daemon will re-check.

* **Query for Hosts:** Remember the `Hosts` tab? It allows you to query list of hosts dynamically. You can use the same query statement here to fetch the list of hosts to check.

* If you want to have a static list of hosts, fill in the text area instead.

* **Check Expressions:** This is the real meat of the check system. See the subsection below.


# What is Check Expression?

Check expression allows user to be very specific when defining alert criteria.

Each expression perform check on a **group of hosts** and raise alert when at minimum `N` hosts are affected.

The purpose of this feature is to minimize noise. For example: often times, it is ok for a system to have one dead host out of twenty.

There are 7 different types of check: *(Note: You can only check charted metrics)*

* **Raw host data:** This type should be familiar to Nagios users, it checks the current, absolute value of a particular metric.

* **Relative host data:** Often times, absolute value checks are not enough. If you need to check changes-over-time, this check is for you.

* **Log data:** Use this check if you need to check the count of loglines. Remember the `Logs` tab? You can use the same query statement to count loglines.

* **Ping, SSH, HTTP/S:** These are active checks *(check the hosts directly)* to make sure hosts are truly well-behaving.


# Multiple Check Expressions Under One Check?

You can have multiple expressions by pressing `+` button. The purpose of this feature is to minimize noise *(We all know how frustrating it is to receive false alerts)*.

The more specific your check is, the less noisy it is. You can use `AND` or `OR` for combining expressions.


# What is a trigger?

Once a check is created, it doesn't actually do anything for you yet. To have a meaningful, actionable alerts, you need to create triggers.

There are 4 different triggers:

* Do Nothing

* Send Email

* Send SMS

* Send PagerDuty


# Can I create Pager Escalation Rules?

A la PagerDuty?

Yes. The nifty thing about triggers is that you can create multiple of them for one check.

To create an escalation policy, just stack multiple triggers. For example:

* Between 1-3 violations, do nothing.

* Between 4-6 violations, email tier1@example.com

* Between 7-10 violations, email tier2@example.com

* Between 11-20 violations, email svp@example.com
