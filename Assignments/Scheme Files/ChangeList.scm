(define (changeList mylist)
        (if (list? mylist)
            (redolist mylist)
            'list-error))

(define (redolist mylist)
         (let ((removedList (remove* '(0 -1 1) mylist)))
         (let recur ((list1 removedList))
              (cond
              ((null? list1) '())
              ((negative? (car list1))  (let (( b(/ 1 (* (car list1) -1)))) (append (list b) (recur (cdr list1)))))
              ((> (car list1) 1) (let ((c (* (car list1) 10))) (append (list c) (recur (cdr list1)))))
              ((null? list1) '())))
          ))  
