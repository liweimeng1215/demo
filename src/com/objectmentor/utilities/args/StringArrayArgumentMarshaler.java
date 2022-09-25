package com.objectmentor.utilities.args;

import java.util.*;

public class StringArrayArgumentMarshaler implements ArgumentMarshaler {
    private ArrayList<Object> stringListValue;

    @Override
    public void set(Iterator<String> currentArgument) throws ArgsException {
        stringListValue = new ArrayList<>();
        while (currentArgument.hasNext()) {
            String argString = currentArgument.next();
            if (argString.startsWith("-")) {
                ListIterator<String> listIter = (ListIterator<String>) currentArgument;
                listIter.previous();
                break;
            } else {
                stringListValue.add(argString);
            }
        }
    }

    public static String[] getValue(ArgumentMarshaler am) {
        if (am != null && am instanceof StringArrayArgumentMarshaler) {
            String[] result = ((StringArrayArgumentMarshaler) am).stringListValue.toArray(new String[0]);
            return result;
        }
        return new String[0];
    }
}