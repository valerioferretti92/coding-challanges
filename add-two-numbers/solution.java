class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1.val == 0 && l1.next == null) {
            return l2;
        }
        if (l2.val == 0 && l2.next == null) {
            return l1;
        }

        ListNode tmp1 = l1, tmp2 = l2, result = new ListNode();
        ListNode current = result;
        int carry = 0;
        while (true) {
            if (tmp1 != null && tmp2 != null) {
                current.val = (tmp1.val + tmp2.val + carry) % 10;
                carry = (tmp1.val + tmp2.val + carry - current.val) / 10;
            } else if (tmp1 != null && tmp2 == null) {
                current.val = (tmp1.val + carry) % 10;
                carry = (tmp1.val + carry - current.val) / 10;
            } else if (tmp1 == null && tmp2 != null) {
                current.val = (tmp2.val + carry) % 10;
                carry = (tmp2.val + carry - current.val) / 10;
            } else if (tmp1 == null && tmp2 == null && carry != 0) {
                current.val = carry % 10;
                carry = (carry - current.val) / 10;
            }

            if (tmp1 != null) {
                tmp1 = tmp1.next;
            }
            if (tmp2 != null) {
                tmp2 = tmp2.next;
            }
            if (tmp1 != null || tmp2 != null || carry != 0) {
                current.next = new ListNode();
                current = current.next;
                continue;
            }
            return result;
        }
    }
 
    public class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}
